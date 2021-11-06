package handlers

import (
	"encoding/json"
	"github.com/nir007/blog/internal/services"
	"github.com/nir007/blog/internal/utils"
	"net/http"
	"strconv"
)

type SeriesHandler struct {
	seriesService *services.SeriesService
	userService *services.UserService
	seriesArticleService *services.SeriesArticleService
}

func NewSeriesHandler(
	seriesService *services.SeriesService,
	userService *services.UserService,
	seriesArticleService *services.SeriesArticleService,
) *SeriesHandler {
	return &SeriesHandler{
		seriesService: seriesService,
		userService: userService,
		seriesArticleService: seriesArticleService,
	}
}

// CreateSeriesAction creates series handler
func (h *SeriesHandler) CreateSeries(w http.ResponseWriter, r *http.Request) {

	uid, err := r.Cookie(cookieNameId)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.userService.GetLoggedUser(uid.Value)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	series := &services.Series{}
	err = json.NewDecoder(r.Body).Decode(&series)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, err)
		return
	}

	series.AuthorId = user.Id
	id, err := h.seriesService.Create(series)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, id)
}

func (h *SeriesHandler) GetOneSeries(w http.ResponseWriter, r *http.Request) {

	seriesId := r.FormValue("series_id")
	sId, err := strconv.ParseInt(seriesId, 10, 32)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, err)
		return
	}

	series, err := h.seriesService.One(sId)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, series)
}

func (h *SeriesHandler) DeleteSeries(w http.ResponseWriter, r *http.Request) {

	seriesId := r.FormValue("id")
	sId, _ := strconv.ParseInt(seriesId, 10, 32)

	uid, err := r.Cookie(cookieNameId)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.userService.GetLoggedUser(uid.Value)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	_, err = h.seriesService.Delete(sId, int64(user.Id))

	utils.WithJSON(w, http.StatusOK, nil)
}

func (h *SeriesHandler) UpdateSeries(w http.ResponseWriter, r *http.Request) {

	series := new(services.Series)

	err := json.NewDecoder(r.Body).Decode(&series)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, err)
		return
	}

	_, err = h.seriesService.Update(series)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	err = h.seriesArticleService.Rebuild(series.Id, series.Articles)
	if err == nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, nil)
}

func (h *SeriesHandler) GetUserSeries(w http.ResponseWriter, r *http.Request) {

	authorId := r.FormValue("author_id")
	limit := r.FormValue("limit")
	offset := r.FormValue("offset")

	aId, _ := strconv.ParseInt(authorId, 10, 32)
	perPage, _ := strconv.ParseInt(limit, 10, 32)
	skip, _ := strconv.ParseInt(offset, 10, 32)

	rows, err := h.seriesService.Read(int32(aId), perPage, skip)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, rows)
}

