package topsort

import (
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/miquelbarba/topsort_api/controllers"
	"github.com/steinfletcher/apitest"
)

var server *gin.Engine

func setUp() {
	server = gin.Default()
	server.GET("/calculate", controllers.Calculate)
}

func TestCalculate(t *testing.T) {
	apitest.New().
		Handler(server).
		Get("/calculate").
		Query("path", "IND,EWR").
		Query("path", "SFO,ATL").
		Query("path", "GSO,IND").
		Query("path", "ATL,GSO").
		Expect(t).
		Body(`{"result":["SFO","EWR"]}`).
		Status(http.StatusOK).
		End()
}

func TestCalculateTrimSpaces(t *testing.T) {
	apitest.New().
		Handler(server).
		Get("/calculate").
		Query("path", "SFO, ATL").
		Query("path", "ATL,GSO ").
		Query("path", "GSO, EWR ").
		Expect(t).
		Body(`{"result":["SFO","EWR"]}`).
		Status(http.StatusOK).
		End()
}

func TestCalculateInvalidParam(t *testing.T) {
	apitest.New().
		Handler(server).
		Get("/calculate").
		Query("path", "IND,EWR").
		Query("path", "SFO").
		Query("path", "GSO,IND").
		Query("path", "ATL,GSO").
		Expect(t).
		Body(`{"message": "invalid parameter SFO"}`).
		Status(http.StatusBadRequest).
		End()
}

func TestCalculateInvalidPath(t *testing.T) {
	apitest.New().
		Handler(server).
		Get("/calculate").
		Query("path", "IND,EWR").
		Query("path", "GSO,ATL").
		Expect(t).
		Body(`{"message": "graph with separated paths"}`).
		Status(http.StatusBadRequest).
		End()
}

func TestMain(m *testing.M) {
	setUp()

	retCode := m.Run()

	os.Exit(retCode)
}
