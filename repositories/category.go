package repositories

import (
	"database/sql"
	"golang-beginner-chap24/collections"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetAll() ([]collections.Category, error) {
	rows, err := r.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []collections.Category
	for rows.Next() {
		var category collections.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
