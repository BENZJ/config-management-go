package fileitem

import (
	"config-management-go/controller/request"
	"config-management-go/models/file"
	fileitem "config-management-go/models/file_item"
	"config-management-go/utils"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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
	fileId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var items []fileitem.FieldItem
	//查询所有文件项目
	s.fileItemRepo.ListAll(fileId, &items)
	//查询文件内容
	var file file.File
	s.fileRepo.ListById(fileId, &file)
	//生成预览文件
	result := generateFile(items, file)
	c.JSON(http.StatusOK, utils.NewResponseData(http.StatusOK, "success", result))
	return

}

//生成预览文件方法
func generateFile(items []fileitem.FieldItem, file file.File) string {
	var result strings.Builder

	// 解析文件名后缀
	fileName := strings.TrimSuffix(file.FileName, filepath.Ext(file.FileName))
	// 确定注释符号
	commentSymbol := "#"
	if filepath.Ext(file.FileName) == ".sql" {
		commentSymbol = "--"
	}

	// 遍历每个 FieldItem
	for _, item := range items {
		// 拼接 FieldItem 的基本信息
		content := fmt.Sprintf("%s Created By: %s\n", commentSymbol, item.CreatedBy)
		content += fmt.Sprintf("%s Created At: %s\n", commentSymbol, item.CreatedAt.Format("2006-01-02 15:04:05"))
		if !item.UpdatedAt.IsZero() {
			content += fmt.Sprintf("%s Updated By: %s\n", commentSymbol, item.UpdatedBy)
			content += fmt.Sprintf("%s Updated At: %s\n", commentSymbol, item.UpdatedAt.Format("2006-01-02 15:04:05"))
		}

		// 拼接到结果中
		result.WriteString(fmt.Sprintf("%s------ %s %s Item %d ------\n", commentSymbol, fileName, strings.ToUpper(filepath.Ext(file.FileName)[1:]), item.ID))
		result.WriteString(content)
		result.WriteString(fmt.Sprintf("%s------ End of Item ------\n\n", commentSymbol))
	}

	return result.String()
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
