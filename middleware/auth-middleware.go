package middleware

import (
	"golang-beginner-chap24/database"
	"golang-beginner-chap24/repositories"
	"golang-beginner-chap24/services"
	"golang-beginner-chap24/utils"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Middleware struct {
	Log *zap.Logger
}

func NewMiddleware(log *zap.Logger) *Middleware {
	return &Middleware{
		Log: log,
	}
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Start timing the request

		cookie, err := r.Cookie("token")
		if err != nil || cookie == nil {
			m.Log.Info("Cookie is nil or error retrieving cookie",
				zap.String("url", r.URL.String()),
				zap.String("method", r.Method),
				zap.Duration("duration", time.Since(start)))
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return // Return early after redirecting
		}

		db := database.NewPostgresDB()
		repo := repositories.NewAdminRepository(db) // Pass logger to the repository
		service := services.NewAdminService(*repo)
		token, err := service.GetAdminByToken(cookie.Value)
		if err != nil {
			m.Log.Error("Failed to get admin by token", zap.Error(err))
			utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		if token == "" || token != cookie.Value {
			m.Log.Info("Unauthorized access attempt",
				zap.String("url", r.URL.String()),
				zap.String("method", r.Method),
				zap.Duration("duration", time.Since(start)))
			utils.RespondWithJSON(w, http.StatusUnauthorized, "Unauthorized", nil)
			return // Return early after unauthorized response
		}

		// Log successful authorization
		m.Log.Info("Authorized access",
			zap.String("url", r.URL.String()),
			zap.String("method", r.Method),
			zap.Duration("duration", time.Since(start)))

		next.ServeHTTP(w, r) // Call the next handler
	})
}
