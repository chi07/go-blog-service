package service_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/shinichi2510/go-blog-service/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestDeleteArticleService_Delete(t *testing.T) {
	ctx := context.Background()
	articleID := uint64(1)
	article := buildMockArticle("Golang Vietnam", "Go Forum", []byte("Content goes here"))
	article.ID = articleID

	t.Run("DeleteArticleFailed", func(t *testing.T) {
		articleMockService := new(ArticleMockService)
		articleMockService.On("Delete", ctx, articleID).Return(errors.New("cannot delete article"))

		s := service.NewDeleteArticleService(articleMockService)
		err := s.Delete(ctx, articleID)

		articleMockService.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("DeleteArticleSuccess", func(t *testing.T) {
		articleMockService := new(ArticleMockService)
		articleMockService.On("Delete", ctx, articleID).Return(nil)

		s := service.NewDeleteArticleService(articleMockService)
		err := s.Delete(ctx, articleID)

		articleMockService.AssertExpectations(t)
		assert.NoError(t, err)
	})
}
