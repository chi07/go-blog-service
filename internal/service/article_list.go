package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shinichi2510/go-blog-service/internal/model"
)

type ListArticle struct {
	articleListReader ArticleListReader
}

func (s *ListArticle) NewListArticle(ctx context.Context, paginator *model.Paginator) ([]*model.Article, error) {
	articles, err := s.articleListReader.List(ctx, paginator)
	if err != nil {
		return nil, errors.Wrap(err, "cannot list article.")
	}
	return articles, nil
}
