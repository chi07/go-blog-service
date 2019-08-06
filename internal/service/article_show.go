package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shinichi2510/go-blog-service/internal/model"
)

type GetArticleService struct {
	articleReader ArticleReader
}

func NewGetArticleService(articleReader ArticleReader) *GetArticleService {
	return &GetArticleService{articleReader: articleReader}
}

func (service *GetArticleService) Get(ctx context.Context, id uint64) (*model.Article, error) {
	article, err := service.articleReader.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get article")
	}
	return article, nil
}
