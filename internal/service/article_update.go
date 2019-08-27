package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/shinichi2510/go-blog-service/internal/model"
)

type UpdateArticleService struct {
	articleSaver ArticleSaver
}

func NewUpdateArticleService(articleSaver ArticleSaver) *UpdateArticleService {
	return &UpdateArticleService{articleSaver: articleSaver}
}

func (s *UpdateArticleService) Update(ctx context.Context, articleID uint64, req *model.ArticleRequest) error {
	article, err := s.articleSaver.Get(ctx, articleID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("cannot found article with ID: %d", articleID))
	}
	article.Title = req.Title
	article.Description = req.Description
	article.Content = req.Content
	article.UpdatedAt = time.Now()

	err = s.articleSaver.Update(ctx, article)
	return errors.Wrap(err, "cannot update article")
}
