package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/miquelbarba/topsort_api/controllers"
)

func main() {
	server := gin.Default()
	server.GET("/topsort", controllers.Topsort)

	log.Fatal(server.Run(":8080"))
}
