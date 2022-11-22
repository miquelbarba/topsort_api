package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miquelbarba/topsort_api/lib"
)

func Calculate(ctx *gin.Context) {
	path := ctx.QueryArray("path")

	edges := make([][]string, len(path))
	for i, item := range path {
		edges[i] = strings.Split(item, ",")
	}

	sort := lib.TopologicalSort(edges)

	result := [2]string{sort[0], sort[len(sort)-1]}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}
