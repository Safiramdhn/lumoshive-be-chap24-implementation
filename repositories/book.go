package repositories

import (
	"database/sql"
	"golang-beginner-chap24/collections"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) Create(bookInput collections.Book) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()
	query := `INSERT INTO books (name, author, price, category_id, price, discount, cover, file)`
	_, err = tx.Exec(query, bookInput.Name, bookInput.Author, bookInput.Price, bookInput.Category.ID, bookInput.Price, bookInput.Discount, bookInput.Cover, bookInput.File)

	if err := tx.Commit(); err != nil {
		return err
	}
	return err
}
