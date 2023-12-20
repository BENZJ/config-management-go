package controller

import (
	"config-management-go/service/file"

	"github.com/gin-gonic/gin"
)

func NewFileController(
	fileService file.Service,
) FileController {
	return FileController{
		fileService: fileService,
	}
}

type FileController struct {
	fileService file.Service
}

// register route
func (c *FileController) Register(app *gin.Engine) error {
	g := app.Group("/file")
	g.POST("/create", c.fileService.CreateFile)
	return nil
}
