package repositories

import (
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

type OrderItemRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewOrderItemRepo(db *sql.DB, log *zap.Logger) *OrderItemRepository {
	return &OrderItemRepository{DB: db, Logger: log}
}

// Get the total count of ratings in the order_items table
func (repo *OrderItemRepository) GetRatingCount() (int, error) {
	var count int
	sqlStatement := "SELECT COUNT(*) FROM order_items"
	repo.Logger.Info("Get rating count", zap.String("repository", "order_items"), zap.String("query", sqlStatement))
	err := repo.DB.QueryRow(sqlStatement).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Retrieve the latest ratings based on a given limit
func (repo *OrderItemRepository) GetLatestRatings(limit int) ([]float64, error) {
	sqlStatement := `SELECT rating FROM order_items ORDER BY order_item_id DESC LIMIT $1`
	repo.Logger.Info("Get rating", zap.String("repository", "order_items"), zap.String("query", sqlStatement))
	rows, err := repo.DB.Query(sqlStatement, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []float64
	for rows.Next() {
		var rating float64
		if err := rows.Scan(&rating); err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}
	return ratings, nil
}

// Calculate the average rating based on the rules provided
func (repo *OrderItemRepository) CalculateAverageRating() (float64, error) {
	count, err := repo.GetRatingCount()
	if err != nil {
		return 0, err
	}

	var limit int
	if count <= 25 {
		limit = 5
	} else if count > 100 {
		limit = 25
	} else {
		limit = count
	}

	ratings, err := repo.GetLatestRatings(limit)
	if err != nil {
		return 0, err
	}
	if len(ratings) == 0 {
		return 0, errors.New("no ratings available")
	}

	var total float64
	for _, rating := range ratings {
		total += rating
	}

	average := total / float64(len(ratings))
	return average, nil
}
