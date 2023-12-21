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
	fileController controller.FileController,
	fileItemController controller.FileItemController,
) ControllerInjector {
	return ControllerInjector{
		IterationController: iterationController,
		FileController:      fileController,
		FileItemController:  fileItemController,
	}
}

type ControllerInjector struct {
	IterationController controller.IterationController
	FileController      controller.FileController
	FileItemController  controller.FileItemController
}
