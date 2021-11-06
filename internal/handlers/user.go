package handlers

import (
	"encoding/json"
	"github.com/nir007/blog/internal/services"
	"github.com/nir007/blog/internal/utils"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	user := services.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	if err = h.userService.Add(&user); err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	cookie := http.Cookie{
		Name:    cookieNameId,
		Value:   user.Uuid,
		Expires: time.Now().Add(360 * 24 * time.Hour),
	}
	http.SetCookie(w, &cookie)

	h.userService.SetLoggedUser(user)
	utils.WithJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) GetPerson(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	personId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	uuidCookie, err := r.Cookie(cookieNameId)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	user := services.User{}
	if user, err = h.userService.One(personId); err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	if uuidCookie.Value == user.Uuid {
		user.IsOwner = true
	} else {
		user.Uuid = ""
		user.Phone = ""
	}

	utils.WithJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) GetPersons(w http.ResponseWriter, r *http.Request) {
	limit := r.FormValue("limit")
	offset := r.FormValue("offset")

	perPage, _ := strconv.ParseInt(limit, 10, 64)
	skip, _ := strconv.ParseInt(offset, 10, 64)


	users, err := h.userService.Get(perPage, skip)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusCreated, users)
}
