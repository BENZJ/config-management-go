package injector

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGinEngine(col ControllerInjector) *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	app := gin.Default()
	app.Use(cors.New(config))
	col.IterationController.Register(app)
	col.FileController.Register(app)
	col.FileItemController.Register(app)
	return app
}
