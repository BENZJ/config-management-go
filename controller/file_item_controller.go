package controller

import (
	fileitem "config-management-go/service/file_item"

	"github.com/gin-gonic/gin"
)

func NewFileItemController(
	fileItemService fileitem.Service,
) FileItemController {
	return FileItemController{
		fileItemService: fileItemService,
	}
}

type FileItemController struct {
	fileItemService fileitem.Service
}

// register route
func (c *FileItemController) Register(app *gin.Engine) error {
	g := app.Group("/fileItem")
	g.POST("/create", c.fileItemService.CreateFileItem)
	return nil
}
