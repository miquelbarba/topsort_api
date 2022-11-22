package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/miquelbarba/topsort_api/controllers"
	"github.com/miquelbarba/topsort_api/routes"
)

var server *gin.Engine
var TopsortController controllers.TopsortController
var TopsortRouteController routes.TopsortRouteController

func init() {
	server = gin.Default()
}

func main() {
	TopsortRouteController.TopsortRoutes(server)

	log.Fatal(server.Run(":8080"))
}
