package handlers

import (
	"errors"
	"github.com/nir007/blog/internal/services"
	"github.com/nir007/blog/internal/utils"
	"net/http"
	"time"
)

type AuthHandler struct {
	userService *services.UserService
	attemptLoginService *services.AttemptLoginService
}

func NewAuthHandler(
	userService *services.UserService,
	attemptLoginService *services.AttemptLoginService,
) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		attemptLoginService: attemptLoginService,
	}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	uid, err := r.Cookie(cookieNameId)

	if err == nil {
		cookie := http.Cookie{
			Name:    cookieNameId,
			Value:   uid.Value,
			Expires: time.Now().Add(-100 * time.Hour),
		}
		http.SetCookie(w, &cookie)
	}
}


func (h *AuthHandler) IsLogged(w http.ResponseWriter, r *http.Request) {
	uid, err := r.Cookie(cookieNameId)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	isLogged := h.userService.IsLogged(uid.Value)

	utils.WithJSON(w, http.StatusOK, isLogged)
}

func (h *AuthHandler) ConfirmPhone(w http.ResponseWriter, r *http.Request) {

	confirmCode := r.FormValue("code")
	uid, err := r.Cookie(cookieNameId)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	confirmed := false
	if len(confirmCode) > 0 {
		user, err := h.userService.GetLoggedUser(uid.Value)
		if err != nil {
			utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
			return
		}

		confirmed, err = h.userService.ConfirmPhone(user, confirmCode)
		if err != nil {
			utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
			return
		}
	}

	utils.WithJSON(w, http.StatusOK, confirmed)
}

func (h *AuthHandler) CheckPhoneNumber(w http.ResponseWriter, r *http.Request) {

	phone := utils.ThoroughlyClearString(r.FormValue("phone"))

	exists, err := h.userService.PhoneNumberExists(phone)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, exists)
}

func (h *AuthHandler) CheckNickName(w http.ResponseWriter, r *http.Request) {

	nickName := r.FormValue("nickname")

	exists, err := h.userService.NickNameExists(nickName); if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, exists)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {

	code := utils.ThoroughlyClearString(r.FormValue("code"))
	phone := utils.ThoroughlyClearString(r.FormValue("phone"))


	_, err := h.attemptLoginService.Last(phone, code)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	user, err := h.userService.OneByPhone(phone)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	h.userService.SetLoggedUser(user)

	cookie := http.Cookie{
		Name:    cookieNameId,
		Value:   user.Uuid,
		Expires: time.Now().Add(360 * 24 * time.Hour),
	}
	http.SetCookie(w, &cookie)

	utils.WithJSON(w, http.StatusOK, user)
}

// GetCodeToLoginAction return code
func (h *AuthHandler) GetCodeToLogin(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetIP()
	phone := utils.ThoroughlyClearString(r.FormValue("phone"))

	if len(phone) < 10 {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, errors.New("phone is invalid"))
		return
	}

	exists, err := h.userService.PhoneNumberExists(phone)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		utils.WithErrWrapJSON(w, http.StatusBadRequest, errors.New("phone is not found"))
		return
	}

	err = h.attemptLoginService.SendCode(
		phone,
		utils.GetConformationCode(),
		ip,
	)
	if err != nil {
		utils.WithErrWrapJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WithJSON(w, http.StatusOK, nil)
}
