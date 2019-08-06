package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/shinichi2510/go-blog-service/internal/model"
)

type UpdateArticleService struct {
	articleSaver ArticleUpdater
}

func NewUpdateArticleService(articleWriter ArticleUpdater) *UpdateArticleService {
	return &UpdateArticleService{articleSaver: articleWriter}
}

func (s *UpdateArticleService) Update(ctx context.Context, article *model.Article) error {
	article.UpdatedAt = time.Now()
	err := s.articleSaver.Update(ctx, article)
	return errors.Wrap(err, "cannot update article")
}
