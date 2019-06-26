package repo_test

import (
	"context"
	"testing"
	"time"

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
	now := time.Now()
	article := &model.Article{
		Title:     "Golang Vietnam",
		Content:   []byte("Golang Vietnam forum discussion"),
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := articleRepo.Create(ctx, article)
	assert.NoError(t, err)
}

func TestArticle_Get(t *testing.T) {
	articleRepo := setupArticleRepo(t)
	ctx := context.Background()
	now := time.Now()
	article := &model.Article{
		ID:        1,
		Title:     "Golang Vietnam",
		Content:   []byte("Golang Vietnam forum discussion"),
		CreatedAt: now,
		UpdatedAt: now,
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
	now := time.Now()
	article := &model.Article{
		ID:        1,
		Title:     "Golang Vietnam",
		Content:   []byte("Golang Vietnam forum discussion"),
		CreatedAt: now,
		UpdatedAt: now,
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
	now := time.Now()
	article := &model.Article{
		ID:        1,
		Title:     "Golang Vietnam",
		Content:   []byte("Golang Vietnam forum discussion"),
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := articleRepo.Create(ctx, article)
	require.NoError(t, err)

	err = articleRepo.Delete(ctx, article.ID)
	assert.NoError(t, err)
}

func TestArticle_List(t *testing.T) {
	articleRepo := setupArticleRepo(t)
	ctx := context.Background()
	now := time.Now()
	articles := []*model.Article{
		{
			Title:     "Golang Vietnam: Why we are should use Go",
			Content:   []byte("Go is simple, fast, and secure"),
			CreatedAt: now,
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Golang Vietnam: Loop",
			Content:   []byte("For loop in Golang ...."),
			CreatedAt: now,
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Golang Vietnam: Interface in Go",
			Content:   []byte("Go Interface content goes here"),
			CreatedAt: now,
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Golang Vietnam: Channel in Go",
			Content:   []byte("Content goes here"),
			CreatedAt: now,
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Golang Vietnam: Concurrency",
			Content:   []byte("Golang Vietnam forum discussion"),
			CreatedAt: now,
			UpdatedAt: time.Now(),
		},
	}
	err := articleRepo.Create(ctx, articles[0])
	require.NoError(t, err)
	err = articleRepo.Create(ctx, articles[1])
	require.NoError(t, err)
	err = articleRepo.Create(ctx, articles[2])
	require.NoError(t, err)
	err = articleRepo.Create(ctx, articles[3])
	require.NoError(t, err)
	err = articleRepo.Create(ctx, articles[4])
	require.NoError(t, err)

	p := &model.Paginator{CurrentPage: 1, Limit: 3, Total: 5, Offset: 0}

	actual, err := articleRepo.List(ctx, p)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(actual))
}
