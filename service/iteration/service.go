package iteration

import (
	"config-management-go/models/iteration"
	"context"
)

func NewService(
	iterationRepo iteration.Repository,
) Service {
	return &service{
		iterationRepo: iterationRepo,
	}
}

type Service interface {
	CreateIteration(ctx context.Context, item iteration.Iteration) (*iteration.Iteration, error)
}

type service struct {
	iterationRepo iteration.Repository
}

func (s *service) CreateIteration(ctx context.Context, item iteration.Iteration) (*iteration.Iteration, error) {
	err := s.iterationRepo.Create(&item)
	return &item, err
}
