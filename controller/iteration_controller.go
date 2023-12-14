package controller

import (
	"config-management-go/service/iteration"

	"github.com/gin-gonic/gin"
)

func NewIterationController(
	iterationService iteration.Service,
) IterationController {
	return IterationController{
		iterationService: iterationService,
	}
}

type IterationController struct {
	iterationService iteration.Service
}

// register route
func (c *IterationController) Register(app *gin.Engine) error {
	g := app.Group("/iteration")
	g.POST("/create", c.iterationService.CreateIteration)
	g.GET("/list", c.iterationService.GetIterationList)
	return nil
}
