package service

import (
	"context"

	"github.com/pkg/errors"
)

type DeleteArticleService struct {
	articleRemover ArticleRemover
}

func NewDeleteArticleService(articleRemover ArticleRemover) *DeleteArticleService {
	return &DeleteArticleService{articleRemover: articleRemover}
}

func (s *DeleteArticleService) Delete(ctx context.Context, id uint64) error {
	err := s.articleRemover.Delete(ctx, id)
	return errors.Wrap(err, "cannot delete article")
}
