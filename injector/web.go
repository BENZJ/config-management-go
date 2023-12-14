package injector

import (
	"github.com/gin-gonic/gin"
)

func InitGinEngine(col ControllerInjector) *gin.Engine {
	app := gin.Default()
	col.IterationController.Register(app)
	return app
}
