package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/miquelbarba/topsort_api/controllers"
)

type TopsortRouteController struct {
	topsortController controllers.TopsortController
}

func NewTopsortController(topsortController controllers.TopsortController) TopsortRouteController {
	return TopsortRouteController{topsortController}
}

func (c *TopsortRouteController) TopsortRoutes(engine *gin.Engine) {
	engine.GET("/topsort", c.topsortController.Topsort)
}
