package handlers

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("views/*.html"))

func LoginViewHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login-view", nil)
}

func LogoutViewHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "logout-view", nil)
}

func DashboardViewHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "dashboard", nil)
}

// func AddBookFormHandler(w http.ResponseWriter, r *http.Request) {
// 	db := database.NewPostgresDB()
// 	categoryRepo := repositories.NewCategoryRepository(db)
// 	categoryService := services.CategoryService(*&categoryRepo)
// 	templates.ExecuteTemplate(w, "add-book-form", nil)
// }
