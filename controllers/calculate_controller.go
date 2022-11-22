package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miquelbarba/topsort_api/lib"
)

func Calculate(ctx *gin.Context) {
	path := ctx.QueryArray("path")
	if len(path) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "missing parameter path"})
		return
	}

	edges, err := extractParams(path)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	sort, errSort := lib.TopologicalSort(edges)
	if errSort != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errSort.Error()})
		return
	}

	result := [2]string{sort[0], sort[len(sort)-1]}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}

const numItems = 2

func extractParams(path []string) ([][]string, error) {
	edges := make([][]string, len(path))

	for i, item := range path {
		aux := strings.Split(item, ",")

		if len(aux) != numItems {
			return edges, fmt.Errorf("invalid parameter %s", item)
		}

		edges[i] = []string{strings.TrimSpace(aux[0]), strings.TrimSpace(aux[1])}
	}

	return edges, nil
}
