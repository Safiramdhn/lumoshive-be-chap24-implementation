package middleware

import (
	"golang-beginner-chap24/database"
	"golang-beginner-chap24/repositories"
	"golang-beginner-chap24/services"
	"golang-beginner-chap24/utils"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := database.NewPostgresDB()

		cookie, err := r.Cookie("token")
		if err != nil || cookie == nil {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
		}

		repo := repositories.NewAdminRepository(db)
		service := services.NewAdminService(*repo)
		token, err := service.GetAdminByToken(cookie.Value)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		if token == "" || token != cookie.Value {
			// utils.RespondWithJSON(w, http.StatusUnauthorized, "Unauthorized", nil)
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
		}

		// Melanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
