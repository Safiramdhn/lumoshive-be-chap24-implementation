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
	query := `INSERT INTO books (name, author, price, category_id, discount, book_cover, book_file)`
	_, err = tx.Exec(query, bookInput.Name, bookInput.Author, bookInput.Price, bookInput.Category.ID, bookInput.Discount, bookInput.Cover, bookInput.File)

	if err := tx.Commit(); err != nil {
		return err
	}
	return err
}

func (r *BookRepository) GetAll() ([]collections.Book, error) {
	sqlStatement := `SELECT b.id, b.name, c.name, b.author, b.price, b.discount, b.quantity
		FROM books b 
		JOIN categories c ON b.category_id = c.id 
		ORDER BY b.id ASC`
	rows, err := r.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []collections.Book
	for rows.Next() {
		var book collections.Book
		err := rows.Scan(&book.ID, &book.Name, &book.Category.Name, &book.Author, &book.Price, &book.Discount, &book.Quantity)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	// If no books were found, return an empty slice instead of nil
	if len(books) == 0 {
		return []collections.Book{}, nil
	}
	return books, nil
}

func (r *BookRepository) GetForUpdate(id string) (collections.Book, []collections.Category, error) {
	var book collections.Book
	var categories []collections.Category
	sqlStatement := `SELECT name, author, price, category_id, discount, book_cover, book_file FROM books WHERE id = $1`
	err := r.DB.QueryRow(sqlStatement, id).Scan(&book.Name, &book.Author, &book.Price, &book.Category.ID, &book.Discount, &book.Cover, &book.File)
	if err == sql.ErrNoRows {
		return book, categories, nil
	} else if err != nil {
		return book, categories, err
	}

	rows, err := r.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return book, categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var category collections.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return book, categories, err
		}
		categories = append(categories, category)
	}
	return book, categories, nil
}
