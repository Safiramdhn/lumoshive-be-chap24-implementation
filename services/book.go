package services

import (
	"errors"
	"golang-beginner-chap24/collections"
	"golang-beginner-chap24/repositories"
)

type BookService struct {
	BookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) *BookService {
	return &BookService{BookRepo: bookRepo}
}

func (bs *BookService) CreateBook(bookInput collections.Book) error {
	if bookInput.Name == "" {
		return errors.New("book name cannot be empty")
	}
	if bookInput.Author == "" {
		return errors.New("author name cannot be empty")
	}
	if bookInput.Price <= 0.0 {
		return errors.New("price cannot be zero")
	}
	if len(bookInput.Cover) == 0 || len(bookInput.File) == 0 {
		return errors.New("cover and file cannot be empty")
	}

	if bookInput.Category.ID == 0 {
		return errors.New("category cannot be empty")
	}

	if bookInput.Discount <= 0 {
		bookInput.Discount = 0
	}
	return bs.BookRepo.Create(bookInput)
}
