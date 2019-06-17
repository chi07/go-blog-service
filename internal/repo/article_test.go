package repo_test

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shinichi2510/go-blog-service/internal/model"
	"github.com/shinichi2510/go-blog-service/internal/repo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupArticleRepo(t require.TestingT) *repo.Article {
	dbURL := viper.GetString("DB_URL")
	ms, err := sqlx.Open("mysql", dbURL)
	require.NoError(t, err)
	ctx := context.Background()

	_, err = ms.ExecContext(ctx, "TRUNCATE TABLE `articles`")
	require.NoError(t, err)
	return repo.NewArticle(ms)
}

func TestArticle_Create(t *testing.T) {
	articleRepo := setupArticleRepo(t)
	ctx := context.Background()
	article := &model.Article{
		Title:   "Golang Vietnam",
		Content: []byte("Golang Vietnam forum discussion"),
	}
	err := articleRepo.Create(ctx, article)
	assert.NoError(t, err)
}

func TestArticle_Get(t *testing.T) {
	articleRepo := setupArticleRepo(t)
	ctx := context.Background()
	article := &model.Article{
		ID:      1,
		Title:   "Golang Vietnam",
		Content: []byte("Golang Vietnam forum discussion"),
	}
	err := articleRepo.Create(ctx, article)
	require.NoError(t, err)

	actual, err := articleRepo.Get(ctx, article.ID)
	assert.NoError(t, err)
	assert.Equal(t, article.ID, actual.ID)

}

func TestArticle_Update(t *testing.T) {
	articleRepo := setupArticleRepo(t)
	ctx := context.Background()
	article := &model.Article{
		ID:      1,
		Title:   "Golang Vietnam",
		Content: []byte("Golang Vietnam forum discussion"),
	}
	err := articleRepo.Create(ctx, article)
	require.NoError(t, err)

	article.Title = "Golang Vietnam Forum"
	err = articleRepo.Update(ctx, article)
	assert.NoError(t, err)

	actual, err := articleRepo.Get(ctx, article.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Golang Vietnam Forum", actual.Title)
}

func TestArticle_Delete(t *testing.T) {
	articleRepo := setupArticleRepo(t)
	ctx := context.Background()
	article := &model.Article{
		ID:      1,
		Title:   "Golang Vietnam",
		Content: []byte("Golang Vietnam forum discussion"),
	}
	err := articleRepo.Create(ctx, article)
	require.NoError(t, err)

	err = articleRepo.Delete(ctx, article.ID)
	assert.NoError(t, err)
}
