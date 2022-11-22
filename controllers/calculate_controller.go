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

	edges, err := extractParams(path)
	if err == nil {
		sort := lib.TopologicalSort(edges)
		result := [2]string{sort[0], sort[len(sort)-1]}

		ctx.JSON(http.StatusOK, gin.H{"result": result})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

const numItems = 2

func extractParams(path []string) ([][]string, error) {
	edges := make([][]string, len(path))
	for i, item := range path {
		edges[i] = strings.Split(item, ",")

		if len(edges[i]) != numItems {
			return edges, fmt.Errorf("invalid parameter %s", item)
		}
	}

	return edges, nil
}
