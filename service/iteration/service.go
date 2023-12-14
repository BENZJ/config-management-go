package iteration

import (
	"config-management-go/controller/request"
	"config-management-go/models/iteration"
	iterationModel "config-management-go/models/iteration"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewService(
	iterationRepo iteration.Repository,
) Service {
	return &service{
		iterationRepo: iterationRepo,
	}
}

// type Service interface {
// 	CreateIteration(ctx context.Context, item iteration.Iteration) (*iteration.Iteration, error)
// }

type Service interface {
	CreateIteration(c *gin.Context)
	GetIterationList(c *gin.Context)
}

type service struct {
	iterationRepo iteration.Repository
}

// create Iteration
func (s *service) CreateIteration(c *gin.Context) {
	// ctx := c.Request.Context()
	var req request.Iteration
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var item = iterationModel.Iteration{
		Name: req.Name,
	}
	c.JSON(http.StatusOK, gin.H{
		"username": req.Name,
	})
	s.iterationRepo.Create(&item)
	return
}

// get iteration list
func (s *service) GetIterationList(c *gin.Context) {
	// ctx := c.Request.Context()
	var req request.Iteration
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var item = iterationModel.Iteration{
		Name: req.Name,
	}
	c.JSON(http.StatusOK, gin.H{
		"username": req.Name,
	})
	s.iterationRepo.Create(&item)
	return
}
