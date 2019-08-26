package handler_test

import (
	"context"

	"github.com/shinichi2510/go-blog-service/internal/model"
	"github.com/stretchr/testify/mock"
)

type ArticleMockService struct {
	mock.Mock
}

func (s *ArticleMockService) Create(ctx context.Context, a *model.Article) (int64, error) {
	args := s.Called(ctx, a)
	if args.Get(0) == 0 {
		return 0, args.Error(1)
	}
	return args.Get(0).(int64), args.Error(1)
}
