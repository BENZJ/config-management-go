package fileitem

import (
	"config-management-go/controller/request"
	"config-management-go/models/file"
	fileitem "config-management-go/models/file_item"
	"config-management-go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewService(
	fileItemRepo fileitem.Repository,
	fileRepo file.Repository,
) Service {
	return &service{
		fileItemRepo: fileItemRepo,
		fileRepo:     fileRepo,
	}
}

type service struct {
	fileItemRepo fileitem.Repository
	fileRepo     file.Repository
}

// GetFilePreview implements Service.
func (s *service) GetFilePreview(c *gin.Context) {
	idStr := c.Query("fileId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var items []fileitem.FieldItem
	//查询所有文件项目
	s.fileItemRepo.ListAll(id, &items)
	//查询文件内容
	var file file.File
	s.fileRepo.ListById(id, &file)
}

// ModifyFileItem implements Service.
func (s *service) ModifyFileItem(c *gin.Context) {
	var req request.FileItem
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fileItem := fileitem.FieldItem{
		ID:        req.ID,
		Content:   req.Content,
		UpdatedBy: req.UpdatedBy,
	}
	s.fileItemRepo.ModifyItem(&fileItem)
	c.JSON(http.StatusOK, utils.NewResponseData(http.StatusOK, "success", &fileItem))
	return
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
	ModifyFileItem(c *gin.Context)
	GetFilePreview(c *gin.Context)
}
