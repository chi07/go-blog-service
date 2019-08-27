package handler_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
	"github.com/shinichi2510/go-blog-service/internal/http/handler"
	"github.com/shinichi2510/go-blog-service/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateArticleHandler_ServeHTTP(t *testing.T) {
	t.Run("InvalidRequestData", func(t *testing.T) {
		articleService := new(ArticleMockService)
		createArticleHandler := handler.NewCreateArticleHandler(articleService)

		body := bytes.NewBufferString(`abc-xyz`)
		r := httptest.NewRequest("POST", "/articles", body)
		r.Header.Set("Content-Type", "application/json")

		ctx := context.Background()
		r = r.WithContext(ctx)
		w := httptest.NewRecorder()

		createArticleHandler.ServeHTTP(w, r)
		articleService.AssertExpectations(t)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectedRes := `{"error":{"code":400,"message":"invalid character 'a' looking for beginning of value"}}`
		assert.JSONEq(t, expectedRes, w.Body.String())
	})

	t.Run("MissingParamsFields", func(t *testing.T) {
		articleService := new(ArticleMockService)
		createArticleHandler := handler.NewCreateArticleHandler(articleService)

		body := bytes.NewBufferString(`{"abc":"xyz"}`)
		r := httptest.NewRequest("POST", "/articles", body)
		r.Header.Set("Content-Type", "application/json")

		ctx := context.Background()
		r = r.WithContext(ctx)
		w := httptest.NewRecorder()

		createArticleHandler.ServeHTTP(w, r)
		articleService.AssertExpectations(t)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		bodyRes := `{"error":{"code":422,"message":"missing 'params' field: validation error"}}`
		assert.JSONEq(t, bodyRes, w.Body.String())
	})

	t.Run("MissingRequiredParams", func(t *testing.T) {
		articleService := new(ArticleMockService)
		createArticleHandler := handler.NewCreateArticleHandler(articleService)

		body := bytes.NewBufferString(`{"params":{"title": "golang programming language"}}`)
		r := httptest.NewRequest("POST", "/articles", body)
		r.Header.Set("Content-Type", "application/json")

		ctx := context.Background()
		r = r.WithContext(ctx)
		w := httptest.NewRecorder()

		createArticleHandler.ServeHTTP(w, r)
		articleService.AssertExpectations(t)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		bodyRes := `{"error":{"message":"missing content: validation error", "code":422}}`
		assert.JSONEq(t, bodyRes, w.Body.String())
	})

	t.Run("SaveFailedToDatabase", func(t *testing.T) {
		articleService := new(ArticleMockService)
		createArticleHandler := handler.NewCreateArticleHandler(articleService)

		body := bytes.NewBufferString(`{"params":{"title": "golang programming language", "description": "awesome go", "content": "content goes here"}}`)
		r := httptest.NewRequest("POST", "/articles", body)
		r.Header.Set("Content-Type", "application/json")

		article := &model.Article{
			Title:       "golang programming language",
			Description: "awesome go",
			Content:     "content goes here",
		}
		ctx := context.Background()
		articleService.On("Create", ctx, article).Return(0, errors.New("cannot save article"))

		r = r.WithContext(ctx)
		w := httptest.NewRecorder()

		createArticleHandler.ServeHTTP(w, r)
		articleService.AssertExpectations(t)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		bodyRes := `{"error":{"message":"cannot save article", "code":500}}`
		assert.JSONEq(t, bodyRes, w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		articleService := new(ArticleMockService)
		createArticleHandler := handler.NewCreateArticleHandler(articleService)

		body := bytes.NewBufferString(`{"params":{"title": "golang programming language", "description": "awesome go", "content": "content goes here"}}`)
		r := httptest.NewRequest("POST", "/articles", body)
		r.Header.Set("Content-Type", "application/json")

		article := &model.Article{
			Title:       "golang programming language",
			Description: "awesome go",
			Content:     "content goes here",
		}
		ctx := context.Background()
		articleService.On("Create", ctx, article).Return(int64(1), nil)

		r = r.WithContext(ctx)
		w := httptest.NewRecorder()

		createArticleHandler.ServeHTTP(w, r)
		articleService.AssertExpectations(t)
		assert.Equal(t, http.StatusCreated, w.Code)
		bodyRes := `{"data":{"articleID":1}}`
		assert.JSONEq(t, bodyRes, w.Body.String())
	})
}
