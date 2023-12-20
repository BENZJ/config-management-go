package file

import (
	"config-management-go/models/file"
	"net/http"

	"github.com/gin-gonic/gin"
)

type service struct {
	fileRepo file.Repository
}

// CreateFile implements Service.
func (*service) CreateFile(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "还未实现"})
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
}
