package fileitem

import (
	"config-management-go/controller/request"
	fileitem "config-management-go/models/file_item"
	"config-management-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewService(
	fileItemRepo fileitem.Repository,
) Service {
	return &service{
		fileItemRepo: fileItemRepo,
	}
}

type service struct {
	fileItemRepo fileitem.Repository
}

// CreateFile implements Service.
func (s *service) CreateFileItem(c *gin.Context) {
	var req request.FileItem
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数解析错误"})
		return
	}
	fileItem := fileitem.FieldItem{
		FileID:      req.FileID,
		IterationID: req.IterationID,
		Content:     req.Content,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
	}
	s.fileItemRepo.Create(&fileItem)
	c.JSON(http.StatusOK, utils.NewResponseData(http.StatusOK, "success", &fileItem))
	return
}

type Service interface {
	CreateFileItem(c *gin.Context)
}
