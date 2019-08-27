package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/shinichi2510/go-blog-service/internal/http/response"
	"github.com/shinichi2510/go-blog-service/internal/model"
)

type UpdateArticleHandler struct {
	articleUpdater ArticleUpdater
}

func NewUpdateArticleHandler(articleUpdater ArticleUpdater) *UpdateArticleHandler {
	return &UpdateArticleHandler{articleUpdater: articleUpdater}
}

type UpdateArticleRequest struct {
	Params *model.ArticleRequest `json:"params"`
}

func (r *UpdateArticleRequest) Bind(req *http.Request) error {
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

func (h *UpdateArticleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	ID, err := strconv.ParseInt(articleID, 10, 64)
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, http.StatusBadRequest, "invalid articleID")
	}

	var req UpdateArticleRequest
	err = render.Bind(r, &req)

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

	err = h.articleUpdater.Update(r.Context(), uint64(ID), req.Params)
	if err != nil {
		response.Error(w, r, http.StatusInternalServerError, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, r, http.StatusOK, nil)
}
