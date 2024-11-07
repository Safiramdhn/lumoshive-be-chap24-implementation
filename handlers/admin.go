package handlers

import (
	"golang-beginner-chap24/services"
	"golang-beginner-chap24/utils"
	"net/http"

	"github.com/google/uuid"
)

type AdminHandler struct {
	AdminService services.AdminService
}

func NewAdminHandler(adminService services.AdminService) *AdminHandler {
	return &AdminHandler{AdminService: adminService}
}

func (ah *AdminHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Invalid request method", nil)
		return
	}

	if err := r.ParseForm(); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request form", nil)

		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	token := uuid.NewString()

	if err := ah.AdminService.LoginAdmin(username, password, token); err != nil {
		utils.RespondWithJSON(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func (ah *AdminHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil || cookie.Value == "" {
		utils.RespondWithJSON(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	// err = ah.AdminService.LogoutAdmin(cookie.Value)
	// if err != nil {
	// 	utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
	// 	return
	// }

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
