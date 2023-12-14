package injector

import (
	"config-management-go/controller"

	"github.com/gin-gonic/gin"
)

func NewApiInjector(
	engine *gin.Engine,
) *ApiInjector {
	return &ApiInjector{
		Engine: engine,
	}
}

type ApiInjector struct {
	Engine *gin.Engine
}

func NewControllerInjector(
	iterationController controller.IterationController,
) ControllerInjector {
	return ControllerInjector{
		IterationController: iterationController,
	}
}

type ControllerInjector struct {
	IterationController controller.IterationController
}
