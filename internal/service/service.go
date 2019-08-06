package service

import (
	"context"

	"github.com/shinichi2510/go-blog-service/internal/model"
)

type ArticleGetter interface {
	Get(ctx context.Context, id uint64) (*model.Article, error)
}

type ArticleLister interface {
	List(ctx context.Context, paginator *model.Paginator) ([]*model.Article, error)
}

type ArticleCreator interface {
	Create(ctx context.Context, id *model.Article) error
}

type ArticleUpdater interface {
	Update(ctx context.Context, article *model.Article) error
}

type ArticleDeleter interface {
	Delete(ctx context.Context, id uint64) error
}
