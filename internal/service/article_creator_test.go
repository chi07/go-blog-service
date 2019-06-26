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

func TestCreateArticleService_Create(t *testing.T) {
	ctx := context.Background()
	article := buildMockArticle("Golang Vietnam", "Go Forum", []byte("Content goes here"))

	t.Run("CreateArticleFailed", func(t *testing.T) {
		articleMockService := new(ArticleMockService)
		articleMockService.On("Create", ctx, mock.Anything).Run(func(args mock.Arguments) {
			article, ok := args.Get(1).(*model.Article)
			assert.True(t, ok)
			assert.NotNil(t, article)
			assert.WithinDuration(t, time.Now(), article.CreatedAt, time.Second)
		}).Return(errors.New("cannot create article"))

		s := service.NewCreateArticleService(articleMockService)
		err := s.Create(ctx, article)

		articleMockService.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("CreateArticleSuccess", func(t *testing.T) {
		articleMockService := new(ArticleMockService)
		articleMockService.On("Create", ctx, mock.Anything).Run(func(args mock.Arguments) {
			article, ok := args.Get(1).(*model.Article)
			assert.True(t, ok)
			assert.NotNil(t, article)
			assert.WithinDuration(t, time.Now(), article.CreatedAt, time.Second)
		}).Return(nil)

		err := articleMockService.Create(ctx, article)
		articleMockService.AssertExpectations(t)
		assert.NoError(t, err)
	})
}
