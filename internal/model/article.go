package model

import (
	"database/sql"
	"time"
)

type Article struct {
	ID              uint64         `db:"id" json:"id"`
	Title           string         `db:"title" json:"title"`
	Description     string         `db:"description" json:"description"`
	Content         string         `db:"content" json:"content"`
	MetaKeyWords    sql.NullString `db:"meta_keywords" json:"metaKeywords"`
	MetaDescription sql.NullString `db:"meta_description" json:"metaDescription"`
	Tags            sql.NullString `db:"-"`
	PublishedTime   time.Time      `db:"published_time"`
	CreatedAt       time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time      `db:"updated_at" json:"updatedAt"`
}

type ArticleRequest struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Content         string `json:"content"`
	Tags            string `json:"tags"`
	MetaKeywords    string `json:"meta_keywords"`
	MetaDescription string `json:"meta_description"`
}

type ListArticleResponse struct {
	Articles   []*Article
	Pagination *Paginator
}
