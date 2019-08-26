package service_test

import (
	"context"

	"github.com/shinichi2510/go-blog-service/internal/model"
	"github.com/stretchr/testify/mock"
)

type ArticleMockService struct {
	mock.Mock
}

func (m *ArticleMockService) Create(ctx context.Context, article *model.Article) (int64, error) {
	args := m.Called(ctx, article)
	if args.Get(0) == 0 {
		return int64(0), args.Error(1)
	}
	return args.Get(0).(int64), nil
}

func (m *ArticleMockService) Update(ctx context.Context, article *model.Article) error {
	args := m.Called(ctx, article)
	return args.Error(0)
}

func (m *ArticleMockService) Get(ctx context.Context, id uint64) (*model.Article, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Article), args.Error(1)
}

func (m *ArticleMockService) List(ctx context.Context, paginator *model.Paginator) ([]*model.Article, error) {
	args := m.Called(ctx, paginator)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Article), args.Error(1)
}

func (m *ArticleMockService) Delete(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func buildMockArticle(title, description string, content string) *model.Article {
	return &model.Article{
		Title:       title,
		Description: description,
		Content:     content,
	}
}
