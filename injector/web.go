package injector

import (
	"config-management-go/models/iteration"

	"github.com/gin-gonic/gin"
)

func InitGinEngine(iteration.Repository) *gin.Engine {
	app := gin.Default()
	return app
}
