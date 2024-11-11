package repositories

import (
	"database/sql"
	"golang-beginner-chap24/collections"
	"log"
)

type PaymentMethodRepository struct {
	DB *sql.DB
}

func NewPaymentMethodRepository(db *sql.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{DB: db}
}

func (r *PaymentMethodRepository) Create(paymentMethodInput collections.PaymentMethod) error {
	tx, err := r.DB.Begin()
	if err != nil {
		log.Printf("Error starting transaction for payment method %v\n", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Printf("Rollback transaction for creating payment method %v\n", p)
			tx.Rollback()
		}
	}()

	sqlStatement := `INSERT INTO payment_methods (name, photo) VALUES ($1, $2)`
	_, err = tx.Exec(sqlStatement, paymentMethodInput.Name, paymentMethodInput.Photo)

	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction for payment method %v\n", err)
		return err
	}
	return err
}

func (r *PaymentMethodRepository) GetAll() ([]collections.PaymentMethod, error) {
	rows, err := r.DB.Query("SELECT id, name, photo FROM payment_methods")
	if err != nil {
		log.Printf("Error getting payment methods %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var paymentMethods []collections.PaymentMethod
	for rows.Next() {
		var paymentMethod collections.PaymentMethod
		err := rows.Scan(&paymentMethod.ID, &paymentMethod.Name, &paymentMethod.Photo)
		if err != nil {
			log.Printf("Error scanning payment method %v\n", err)
			return nil, err
		}
		paymentMethods = append(paymentMethods, paymentMethod)
	}
	return paymentMethods, nil
}

func (r *PaymentMethodRepository) GetByID(id int) (collections.PaymentMethod, error) {
	var paymentMethod collections.PaymentMethod
	sqlStatement := `SELECT id, name, photo FROM payment_methods WHERE id = $1`
	err := r.DB.QueryRow(sqlStatement, id).Scan(&paymentMethod.ID, &paymentMethod.Name, &paymentMethod.Photo)
	if err == sql.ErrNoRows {
		log.Printf("payment method not found ")
		return paymentMethod, nil
	} else if err != nil {
		log.Printf("Error getting payment method %v\n", err)
		return paymentMethod, err
	}
	return paymentMethod, nil
}

func (r *PaymentMethodRepository) Update(id int, paymentMethodInput collections.PaymentMethod) error {
	tx, err := r.DB.Begin()
	if err != nil {
		log.Printf("Error starting transaction for payment method %v\n", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Printf("Rollback transaction for updating payment method %v\n", p)
			tx.Rollback()
		}
	}()
	sqlStatement := `UPDATE payment_methods SET name = $1, photo = $2 WHERE id = $3`
	_, err = tx.Exec(sqlStatement, paymentMethodInput.Name, paymentMethodInput.Photo, id)

	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction for payment method %v\n", err)
		return err
	}
	return err
}

func (r *PaymentMethodRepository) Delete(id int) error {
	tx, err := r.DB.Begin()
	if err != nil {
		log.Printf("Error starting transaction for payment method %v\n", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Printf("Rollback transaction for deleting payment method %v\n", p)
			tx.Rollback()
		}
	}()

	_, err = r.DB.Exec("UPDATE payment_methods SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1", id)
	if err != nil {
		log.Printf("Error deleting payment method %v\n", err)
		return err
	}
	return nil
}
