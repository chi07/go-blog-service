package handler

import (
	"context"
	"errors"

	"github.com/shinichi2510/go-blog-service/internal/model"
)

var (
	ErrValidation = errors.New("validation error")
)

type ArticleWriter interface {
	Create(ctx context.Context, story *model.Article) (int64, error)
}

type ArticleGetter interface {
	Get(ctx context.Context, id uint64) (*model.Article, error)
}

type ArticleUpdater interface {
	Update(ctx context.Context, articleID uint64, req *model.ArticleRequest) error
}

type ArticleDeleter interface {
	Delete(ctx context.Context, id uint64) error
}
