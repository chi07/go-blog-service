package model

import (
	"database/sql"
	"time"
)

type Article struct {
	ID              uint64         `db:"id"`
	Title           string         `db:"title"`
	Description     string         `db:"description"`
	Content         []byte         `db:"content"`
	MetaKeyWords    sql.NullString `db:"meta_keywords"`
	MetaDescription sql.NullString `db:"meta_description"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
}

type ListArticleResponse struct {
	Articles   []*Article
	Pagination *Paginator
}
