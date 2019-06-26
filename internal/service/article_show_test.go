package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/shinichi2510/go-blog-service/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestGetArticle_Get(t *testing.T) {
	ctx := context.Background()
	t.Run("GetArticleSuccessfully", func(t *testing.T) {
		articleID := uint64(12)
		expectedArticle := buildMockArticle("Golang Vietnam", "Forum Golang", []byte("Golang Vietnam"))
		expectedArticle.ID = articleID
		articleMockService := new(ArticleMockService)
		articleMockService.On("Get", ctx, articleID).Return(expectedArticle, nil)

		s := service.NewGetArticleService(articleMockService)
		actual, err := s.Get(ctx, articleID)
		articleMockService.AssertExpectations(t)
		assert.NoError(t, err)
		assert.Equal(t, expectedArticle, actual)
	})
	t.Run("GetArticleFailed", func(t *testing.T) {
		articleID := uint64(12)
		articleMockService := new(ArticleMockService)
		articleMockService.On("Get", ctx, articleID).Return(nil, errors.New("cannot get article"))

		s := service.NewGetArticleService(articleMockService)
		actual, err := s.Get(ctx, articleID)
		articleMockService.AssertExpectations(t)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}
