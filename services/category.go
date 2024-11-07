package services

import (
	"golang-beginner-chap24/collections"
	"golang-beginner-chap24/repositories"
)

type CategoryService struct {
	CategoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: categoryRepo}
}

func (s *CategoryService) GetAllCategories() ([]collections.Category, error) {
	return s.CategoryRepo.GetAll()
}
