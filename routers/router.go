package routers

import (
	"golang-beginner-chap24/database"
	"golang-beginner-chap24/handlers"
	"golang-beginner-chap24/repositories"
	"golang-beginner-chap24/services"
	"golang-beginner-chap24/wire"

	"github.com/go-chi/chi/v5"
)

func NewRouter() chi.Router {
	db := database.NewPostgresDB()
	r := chi.NewRouter()
	orderHandler := wire.InitializeOrderHandler()
	paymentMethodHandler := wire.InitializePaymentMethodHandler()
	auth := wire.IniitalMiddleware()
	orderItemHandler := wire.InitializeOrderItemHandler()

	adminRepo := repositories.NewAdminRepository(db)
	adminService := services.NewAdminService(*adminRepo)
	adminHandler := handlers.NewAdminHandler(*adminService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(*categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(*categoryService)

	bookRepo := repositories.NewBookRepository(db)
	bookService := services.NewBookService(*bookRepo)
	bookHandler := handlers.NewBookHandler(*bookService)

	// orderRepo := repositories.NewOrderRepository(db)
	// orderService := services.NewOrderService(*orderRepo)
	// orderHandler := handlers.NewOrderHandler(*orderService)

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", adminHandler.LoginHandler)
		r.With(auth.AuthMiddleware).Post("/logout", adminHandler.LogoutHandler)

		r.Route("/book", func(r chi.Router) {
			r.With(auth.AuthMiddleware).Post("/add-book", bookHandler.CreateBookHandler)
		})

		r.Route("/order", func(r chi.Router) {
			r.Post("/create", orderHandler.CreateOrderHandler)
			r.With(auth.AuthMiddleware).Get("/rating", orderItemHandler.GetRatingAverageHandler)
		})

		r.Route("/payment-methods", func(r chi.Router) {
			r.Post("/", paymentMethodHandler.CreatePaymentMethodHandler)
			r.With(auth.AuthMiddleware).Get("/", paymentMethodHandler.GetAllPaymentMethodsHandler)
			r.Get("/{id}", paymentMethodHandler.GetPaymentMethodByIdHandler)
			r.Put("/{id}", paymentMethodHandler.UpdatePaymentMethodHandler)
			r.Delete("/{id}", paymentMethodHandler.DeletePaymentMethodHandler)

		})
	})

	r.Get("/login", handlers.LoginViewHandler)
	r.Get("/logout", handlers.LogoutViewHandler)

	r.With(auth.AuthMiddleware).Get("/dashboard", handlers.DashboardViewHandler)
	r.With(auth.AuthMiddleware).Get("/add-book-form", categoryHandler.AddBookFormHandler)
	r.With(auth.AuthMiddleware).Get("/book-list", bookHandler.GetBooksHandler)
	r.With(auth.AuthMiddleware).Get("/edit-book/{id}", bookHandler.EditBookFormHandler)

	return r
}
