package handler_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shinichi2510/go-blog-service/internal/http/handler"
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

	t.Run("Missing params fields", func(t *testing.T) {
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
}
