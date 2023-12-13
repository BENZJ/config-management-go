package injector

import (
	"config-management-go/service/iteration"

	"github.com/gin-gonic/gin"
)

func InitGinEngine(iteration.Service) *gin.Engine {
	app := gin.Default()
	return app
}
