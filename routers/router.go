package routers

import (
	"golang-beginner-chap24/database"
	"golang-beginner-chap24/handlers"
	"golang-beginner-chap24/middleware"
	"golang-beginner-chap24/repositories"
	"golang-beginner-chap24/services"
	// "net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() chi.Router {
	db := database.NewPostgresDB()
	r := chi.NewRouter()

	adminRepo := repositories.NewAdminRepository(db)
	adminService := services.NewAdminService(*adminRepo)
	adminHandler := handlers.NewAdminHandler(*adminService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(*categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(*categoryService)

	bookRepo := repositories.NewBookRepository(db)
	bookService := services.NewBookService(*bookRepo)
	bookHandler := handlers.NewBookHandler(*bookService)

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", adminHandler.LoginHandler)
		r.With(middleware.AuthMiddleware).Post("/logout", adminHandler.LogoutHandler)

		r.Route("/book", func(r chi.Router) {
			r.With(middleware.AuthMiddleware).Post("/add-book", bookHandler.CreateBookHandler)
		})
	})

	r.Get("/login", handlers.LoginViewHandler)
	r.Get("/logout", handlers.LogoutViewHandler)

	r.With(middleware.AuthMiddleware).Get("/dashboard", handlers.DashboardViewHandler)
	r.With(middleware.AuthMiddleware).Get("/add-book-form", categoryHandler.AddBookFormHandler)

	return r
}
