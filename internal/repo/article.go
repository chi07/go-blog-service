package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/shinichi2510/go-blog-service/internal/model"
)

type Article struct {
	db *sqlx.DB
}

func NewArticle(db *sqlx.DB) *Article {
	return &Article{db: db}
}

func (a *Article) Get(ctx context.Context, id uint64) (*model.Article, error) {
	var article model.Article
	err := a.db.GetContext(ctx, &article, "SELECT * FROM `articles` WHERE `id` = ?", id)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get article by ID from DB")
	}
	return &article, nil
}

func (a *Article) Create(ctx context.Context, article *model.Article) error {
	_, err := a.db.ExecContext(ctx, "INSERT INTO `articles` (`title`, `content`, `description`, `meta_keywords`, `meta_description`, `created_at`) VALUES(?, ?, ?, ?, ?, ?)", article.Title, article.Content, article.Description, article.MetaKeyWords, article.MetaDescription, article.CreatedAt)
	return errors.Wrap(err, "cannot save new article to DB")
}

func (a *Article) Update(ctx context.Context, article *model.Article) error {
	_, err := a.db.ExecContext(ctx, "UPDATE `articles` SET `title`=?, `content`=?, `description`=?, `meta_keywords`=?, `meta_description`=?, `updated_at`=? WHERE `id`=?", article.Title, article.Content, article.Description, article.MetaKeyWords, article.Description, article.UpdatedAt, article.ID)
	return errors.Wrap(err, "cannot update article into DB")
}

func (a *Article) Delete(ctx context.Context, id uint64) error {
	_, err := a.db.ExecContext(ctx, "DELETE FROM `articles` WHERE `id`=?", id)
	return errors.Wrap(err, "cannot delete article in DB")
}
