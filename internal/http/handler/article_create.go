package handler

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/shinichi2510/go-blog-service/internal/http/response"
	"github.com/shinichi2510/go-blog-service/internal/model"
)

type CreateArticleHandler struct {
	articleWriter ArticleWriter
}

func NewCreateArticleHandler(articleWriter ArticleWriter) *CreateArticleHandler {
	return &CreateArticleHandler{articleWriter: articleWriter}
}

type ArticleRequest struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Content         string `json:"content"`
	Tags            string `json:"tags"`
	MetaKeywords    string `json:"meta_keywords"`
	MetaDescription string `json:"meta_description"`
}

type CreateArticleRequest struct {
	Params *ArticleRequest `json:"params"`
}

func (r *CreateArticleRequest) Bind(req *http.Request) error {
	validationError := make([]string, 0)
	if r.Params == nil {
		return errors.Wrap(ErrValidation, "missing 'params' field")
	}

	if r.Params.Title == "" {
		validationError = append(validationError, "missing title")
	}
	if r.Params.Content == "" {
		validationError = append(validationError, "missing content")
	}
	if len(validationError) > 0 {
		return errors.Wrap(ErrValidation, strings.Join(validationError, ","))
	}
	return nil
}

func (h *CreateArticleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req CreateArticleRequest
	err := render.Bind(r, &req)

	errCause := errors.Cause(err)
	if errCause == ErrValidation {
		response.Error(w, r, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if errCause == ErrValidation {
		response.Error(w, r, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, err.Error())
		return
	}
	if errCause != nil {
		response.Error(w, r, http.StatusBadRequest, http.StatusBadRequest, err.Error())
		return
	}

	article := &model.Article{
		Title:           req.Params.Title,
		Description:     req.Params.Description,
		Content:         req.Params.Content,
		MetaKeyWords:    sql.NullString{String: req.Params.MetaKeywords, Valid: true},
		MetaDescription: sql.NullString{String: req.Params.MetaDescription, Valid: true},
		Tags:            sql.NullString{String: req.Params.Tags, Valid: true},
	}

	articleID, err := h.articleWriter.Create(r.Context(), article)
	if err != nil {
		response.Error(w, r, http.StatusInternalServerError, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]interface{}{"articleID": articleID}

	response.Success(w, r, http.StatusCreated, data)
}
