package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/shinichi2510/go-blog-service/internal/http/response"
)

type ShowArticleHandler struct {
	articleGetter ArticleGetter
}

func NewShowArticleHandler(articleGetter ArticleGetter) *ShowArticleHandler {
	return &ShowArticleHandler{articleGetter: articleGetter}
}

func (h *ShowArticleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	ID, err := strconv.ParseInt(articleID, 10, 64)
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, http.StatusBadRequest, "articleID is not valid")
		return
	}

	article, err := h.articleGetter.Get(r.Context(), uint64(ID))
	if err != nil {
		response.Error(w, r, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, "cannot get article")
		return
	}

	response.Success(w, r, http.StatusOK, article)
}
