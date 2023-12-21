package file

import (
	"config-management-go/controller/request"
	"config-management-go/models/file"
	"config-management-go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type service struct {
	fileRepo file.Repository
}

// ListByIteration implements Service.
func (s *service) ListByIteration(c *gin.Context) {
	idStr := c.Query("iterationID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid age parameter",
		})
		return
	}
	var fileList []file.File
	s.fileRepo.ListAll(id, &fileList)
	c.JSON(http.StatusOK, utils.NewResponseData(http.StatusOK, "success", &fileList))
	return
}

// CreateFile implements Service.
func (s *service) CreateFile(c *gin.Context) {
	var req request.File
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数解析错误"})
		return
	}
	file := file.File{
		IterationID: req.IterationID,
		FileName:    req.FileName,
	}
	s.fileRepo.Create(&file)
	c.JSON(http.StatusOK, utils.NewResponseData(http.StatusOK, "success", &file))
	return
}

func NewService(
	fileRepo file.Repository,
) Service {
	return &service{
		fileRepo: fileRepo,
	}
}

// type Service interface {
// 	CreateIteration(ctx context.Context, item iteration.Iteration) (*iteration.Iteration, error)
// }

type Service interface {
	CreateFile(c *gin.Context)
	ListByIteration(c *gin.Context)
}
