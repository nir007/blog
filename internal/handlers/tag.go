package handlers

import (
	"github.com/nir007/blog/internal/services"
	"github.com/nir007/blog/internal/utils"
	"net/http"
)

type TagHandler struct {
	articleService *services.ArticleService
}

func NewTagHandler(articleService *services.ArticleService) *TagHandler {
	return &TagHandler{
		articleService: articleService,
	}
}

func (h *TagHandler) GetTags(w http.ResponseWriter, r *http.Request) {

	tags := map[string]string{}
	var err error

	if tags, err = h.articleService.GetTags(); err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
	}

	utils.WithJSON(w, http.StatusOK, tags)
}
