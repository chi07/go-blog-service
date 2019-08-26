package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/shinichi2510/go-blog-service/internal/http/response"
)

type DeleteArticleHandler struct {
	articleDeleter ArticleDeleter
}

func NewDeleteArticleHandler(articleDeleter ArticleDeleter) *DeleteArticleHandler {
	return &DeleteArticleHandler{articleDeleter: articleDeleter}
}

func (h *DeleteArticleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "articleID")
	articleID, ok := strconv.ParseInt(ID, 10, 64)
	if ok != nil {
		response.Error(w, r, http.StatusBadRequest, http.StatusBadRequest, "articleID invalid")
		return
	}
	err := h.articleDeleter.Delete(r.Context(), uint64(articleID))
	if err != nil {
		response.Error(w, r, http.StatusInternalServerError, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, r, http.StatusOK, nil)
}
