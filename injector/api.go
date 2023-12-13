package injector

import (
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
