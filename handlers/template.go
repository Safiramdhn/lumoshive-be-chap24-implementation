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
