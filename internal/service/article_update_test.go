package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/shinichi2510/go-blog-service/internal/model"
	"github.com/shinichi2510/go-blog-service/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateArticleService_Update(t *testing.T) {
	ctx := context.Background()
	article := buildMockArticle("Golang Vietnam", "Go Forum", "Content goes here")
	article.ID = uint64(1)

	t.Run("UpdateArticleFailed", func(t *testing.T) {
		articleMockService := new(ArticleMockService)
		articleMockService.On("Update", ctx, mock.Anything).Run(func(args mock.Arguments) {
			article, ok := args.Get(1).(*model.Article)
			assert.True(t, ok)
			assert.NotNil(t, article)
			assert.WithinDuration(t, time.Now(), article.UpdatedAt, time.Second)
		}).Return(errors.New("cannot update article"))

		s := service.NewUpdateArticleService(articleMockService)
		err := s.Update(ctx, article)

		articleMockService.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("UpdateArticleSuccess", func(t *testing.T) {
		articleMockService := new(ArticleMockService)
		articleMockService.On("Update", ctx, mock.Anything).Run(func(args mock.Arguments) {
			article, ok := args.Get(1).(*model.Article)
			assert.True(t, ok)
			assert.NotNil(t, article)
			assert.WithinDuration(t, time.Now(), article.UpdatedAt, time.Second)
		}).Return(nil)

		s := service.NewUpdateArticleService(articleMockService)
		err := s.Update(ctx, article)

		articleMockService.AssertExpectations(t)
		assert.NoError(t, err)
	})
}
