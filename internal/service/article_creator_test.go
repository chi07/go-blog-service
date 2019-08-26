package service_test

import (
	"context"
	"fmt"
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
	article := buildMockArticle("Golang Vietnam", "Go Forum", "Content goes here")

	t.Run("CreateArticleFailed", func(t *testing.T) {
		articleMockService := new(ArticleMockService)
		articleMockService.On("Create", ctx, mock.Anything).Run(func(args mock.Arguments) {
			article, ok := args.Get(1).(*model.Article)
			assert.True(t, ok)
			assert.NotNil(t, article)
			assert.WithinDuration(t, time.Now(), article.CreatedAt, time.Second)
		}).Return(0, errors.New("cannot create article"))

		s := service.NewCreateArticleService(articleMockService)
		articleID, err := s.Create(ctx, article)

		fmt.Println("articleID: ", articleID)

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
		}).Return(int64(1), nil)

		_, err := articleMockService.Create(ctx, article)
		articleMockService.AssertExpectations(t)
		assert.NoError(t, err)
	})
}
