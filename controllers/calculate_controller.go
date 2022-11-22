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
		sort, errSort := lib.TopologicalSort(edges)

		if errSort == nil {
			result := [2]string{sort[0], sort[len(sort)-1]}

			ctx.JSON(http.StatusOK, gin.H{"result": result})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": errSort.Error()})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
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
