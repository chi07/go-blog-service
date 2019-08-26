package service

import (
	"context"
	"fmt"
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

func (s *CreateArticleService) Create(ctx context.Context, article *model.Article) (int64, error) {
	now := time.Now()
	article.CreatedAt = now
	article.UpdatedAt = now

	articleID, err := s.articleCreator.Create(ctx, article)

	if err != nil || articleID == 0 {
		fmt.Println("yah", err)
		return 0, errors.Wrap(err, "cannot create article")
	}
	return articleID, nil
}
