package topsort

import (
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/miquelbarba/topsort_api/routes"
	"github.com/steinfletcher/apitest"
)

var TopsortRouteController routes.TopsortRouteController
var server *gin.Engine

func setUp() {
	server = gin.Default()
	TopsortRouteController.TopsortRoutes(server)
}

func TestCalculateRoute(t *testing.T) {
	apitest.New().
		Handler(server).
		Get("/topsort").
		Query("path", "IND,EWR").
		Query("path", "SFO,ATL").
		Query("path", "GSO,IND").
		Query("path", "ATL,GSO").
		Expect(t).
		Body(`{"result":["SFO","EWR"]}`).
		Status(http.StatusOK).
		End()
}

func TestMain(m *testing.M) {
	setUp()

	retCode := m.Run()

	os.Exit(retCode)
}
