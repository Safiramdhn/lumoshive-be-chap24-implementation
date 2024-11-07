package handlers

import (
	"golang-beginner-chap24/services"
	"golang-beginner-chap24/utils"
	"net/http"
)

type CategoryHandler struct {
	CategoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: categoryService}
}

func (h *CategoryHandler) AddBookFormHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := h.CategoryService.GetAllCategories()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	templates.ExecuteTemplate(w, "add-book-form", map[string]interface{}{
		"Categories": categories,
	})
}
