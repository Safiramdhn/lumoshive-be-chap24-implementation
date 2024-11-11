package repositories

import (
	"database/sql"
	"golang-beginner-chap24/collections"
	"log"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// helper function to handle errors with logging and transaction rollback
func handleError(tx *sql.Tx, err error, msg string) error {
	if err != nil {
		log.Printf("%s: %v\n", msg, err)
		tx.Rollback()
	}
	return err
}

func (r *OrderRepository) Create(orderInput collections.Order) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return handleError(tx, err, "Failed to start transaction")
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	// Insert customer
	customerStatement := `INSERT INTO customers (customer_name, customer_phone) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRow(customerStatement, orderInput.CustomerName, orderInput.CustomerPhone).Scan(&orderInput.CustomerID)
	if err = handleError(tx, err, "Error inserting customer"); err != nil {
		return err
	}

	// Insert address
	addressStatement := `INSERT INTO addresses (customer_id, street, city, postal_code, country) VALUES ($1, $2, $3, $4, $5) RETURNING address_id`
	err = tx.QueryRow(addressStatement, orderInput.CustomerID, orderInput.ShippingAddress.Street, orderInput.ShippingAddress.City, orderInput.ShippingAddress.PostalCode, orderInput.ShippingAddress.Country).Scan(&orderInput.ShippingAddress.ID)
	if err = handleError(tx, err, "Error creating address"); err != nil {
		return err
	}

	// Insert order
	orderStatement := `INSERT INTO orders (customer_id, payment_method, shipping_address_id) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(orderStatement, orderInput.CustomerID, orderInput.PaymentMethod, orderInput.ShippingAddress.ID).Scan(&orderInput.ID)
	if err = handleError(tx, err, "Error inserting order"); err != nil {
		return err
	}

	// Process order items
	for i, book := range orderInput.OrderItems {
		var bookFound collections.Book
		bookStatement := `SELECT price, discount, quantity FROM books WHERE id = $1`
		err = tx.QueryRow(bookStatement, book.BookID).Scan(&bookFound.Price, &bookFound.Discount, &bookFound.Quantity)
		if err = handleError(tx, err, "Error retrieving book details"); err != nil {
			return err
		}

		// Calculate subtotal and update book quantity
		orderInput.OrderItems[i].Subtotal = (bookFound.Price - (bookFound.Price * float64(bookFound.Discount/100))) * float64(book.Quantity)
		bookFound.Quantity -= book.Quantity

		// Update book quantity
		updateBookStatement := `UPDATE books SET quantity = $1 WHERE id = $2`
		_, err = tx.Exec(updateBookStatement, bookFound.Quantity, book.BookID)
		if err = handleError(tx, err, "Error updating book quantity"); err != nil {
			return err
		}

		// Insert order item
		orderItemStatement := `INSERT INTO order_items (order_id, book_id, quantity, subtotal) VALUES ($1, $2, $3, $4)`
		_, err = tx.Exec(orderItemStatement, orderInput.ID, book.BookID, book.Quantity, orderInput.OrderItems[i].Subtotal)
		if err = handleError(tx, err, "Error inserting order item"); err != nil {
			return err
		}

		// Update totals
		orderInput.TotalAmount += bookFound.Price
		orderInput.FinalAmount += orderInput.OrderItems[i].Subtotal
	}

	// Update order with total and final amount
	updateOrderStatement := `UPDATE orders SET total_amount = $1, final_amount = $2 WHERE id = $3`
	_, err = tx.Exec(updateOrderStatement, orderInput.TotalAmount, orderInput.FinalAmount, orderInput.ID)
	if err = handleError(tx, err, "Error updating order totals"); err != nil {
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v\n", err)
		return err
	}
	return nil
}
