package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/shinichi2510/go-blog-service/internal/model"
)

type CreateArticleService struct {
	articleCreator ArticleCreator
}

func NewCreateArticleService(articleCreator ArticleCreator) *CreateArticleService {
	return &CreateArticleService{articleCreator: articleCreator}
}

func (s *CreateArticleService) Create(ctx context.Context, article *model.Article) error {
	now := time.Now()
	article.CreatedAt = now
	article.UpdatedAt = now

	err := s.articleCreator.Create(ctx, article)
	return errors.Wrap(err, "cannot create article")
}
