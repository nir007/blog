package handlers

import (
	"encoding/json"
	"errors"
	"github.com/nir007/blog/internal/services"
	"github.com/nir007/blog/internal/utils"
	"net/http"
	"strconv"
	"time"
)

type ArticleHandler struct {
	articleService *services.ArticleService
	userService *services.UserService
}

func NewArticleHandler(articleService *services.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		articleService: articleService,
	}
}

type GetArticleResponse struct {
	Id          int32             `json:"id"`
	SeriesId    int32             `json:"series_id"`
	AuthorId    int32             `json:"author_id"`
	Title       string            `json:"title"`
	Text        string            `json:"text"`
	CreatedAt   time.Time         `json:"created"`
	Tags        map[string]string `json:"tags"`
	Published   rune              `json:"published"`
	IsOwner     bool              `json:"is_owner"`
	Order       int32             `json:"order"`
	PrevArticle map[string]string `json:"prev_article"`
	NextArticle map[string]string `json:"next_article"`
}

func (GetArticleResponse) toGetArticleResponse(in services.Article) *GetArticleResponse {
	return &GetArticleResponse{
		Id:          in.Id,
		SeriesId:    in.SeriesId,
		AuthorId:    in.AuthorId,
		Title:       in.Title,
		Text:        in.Text,
		CreatedAt:   in.CreatedAt,
		Tags:        in.Tags,
		Published:   in.Published,
		IsOwner:     in.IsOwner,
		Order:       in.Order,
		PrevArticle: in.PrevArticle,
		NextArticle: in.NextArticle,
	}
}

func (h *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	resp := &GetArticleResponse{}
	articleId, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, err)
		return
	}

	article, err := h.articleService.One(articleId)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, resp.toGetArticleResponse(article))
}

func (h *ArticleHandler) RemoveArticle(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	uid, _ := r.Cookie(cookieNameId)

	if uid == nil || id == "" {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, errors.New("invalid uid or id params"))
		return
	}

	articleId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, errors.New("articleId"))
		return
	}

	user, err := h.userService.GetLoggedUser(uid.Value)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	_, err = h.articleService.Remove(articleId, user.Id)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, nil)
}

func (h *ArticleHandler) AddArticle(w http.ResponseWriter, r *http.Request) {
	uid, _ := r.Cookie(cookieNameId)
	if uid == nil {
		utils.WithErrWrapJSON(w, http.StatusUnauthorized, errors.New("need auth"))
		return
	}

	if !h.userService.IsLogged(uid.Value) {
		utils.WithErrWrapJSON(w, http.StatusUnauthorized, errors.New("need auth"))
		return
	}

	article := &services.Article{}

	err := json.NewDecoder(r.Body).Decode(article)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	authorId, err := strconv.ParseInt(uid.Value, 10, 32)

	newArticle := &services.Article{
		AuthorId:  int32(authorId),
		Title:     article.Title,
		Text:      article.Text,
		Tags:      article.Tags,
		CreatedAt: time.Now(),
		Published: article.Published,
	}

	err = h.articleService.Add(newArticle)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, newArticle)
}

func (h *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {

	article := &services.Article{}

	uid, err := r.Cookie(cookieNameId)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	if uid == nil {
		utils.WithErrWrapJSON(w, http.StatusUnauthorized, errors.New("need auth"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.userService.GetLoggedUser(uid.Value)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	if user.Id != article.AuthorId {
		utils.WithErrWrapJSON(w, http.StatusForbidden, err)
		return
	}

	_, err = h.articleService.Update(article)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, nil)
}

func (h *ArticleHandler) GetPublishedArticles(w http.ResponseWriter, r *http.Request) {

	uid, err := r.Cookie(cookieNameId)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	if uid == nil {
		utils.WithErrWrapJSON(w, http.StatusUnauthorized, errors.New("need auth"))
		return
	}

	user, err := h.userService.GetLoggedUser(uid.Value)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	limit := r.FormValue("limit")
	offset := r.FormValue("offset")
	perPage, _ := strconv.ParseInt(limit, 10, 64)
	skip, _ := strconv.ParseInt(offset, 10, 64)

	articles, err := h.articleService.GetPublished(int64(user.Id), perPage, skip)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, articles)
}

func (h *ArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {

	authorId := r.FormValue("author_id")
	tag := r.FormValue("tag")
	limit := r.FormValue("limit")
	offset := r.FormValue("offset")
	showPublished := r.FormValue("show_published")

	aId, _ := strconv.ParseInt(authorId, 10, 64)
	perPage, _ := strconv.ParseInt(limit, 10, 64)
	skip, _ := strconv.ParseInt(offset, 10, 64)
	sPublished, _ := strconv.ParseInt(showPublished, 10, 64)

	articles, err := h.articleService.Get(sPublished, aId, perPage, skip, tag)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	uid, err := r.Cookie(cookieNameId)

	if aId > 0 && uid != nil && err == nil {

		user, err := h.userService.GetLoggedUser(uid.Value)
		if err != nil {
			utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
			return
		}

		for k, v := range articles {
			if user.Id == v.AuthorId {
				articles[k].IsOwner = true
			}
		}
	}

	utils.WithJSON(w, http.StatusOK, articles)
}
