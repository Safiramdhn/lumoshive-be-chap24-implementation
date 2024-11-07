package repositories

import "database/sql"

type AdminRepository struct {
	DB *sql.DB
}

func NewAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (r *AdminRepository) Login(username, password, token string) error {
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

	var userId int
	sqlStatement := `SELECT id FROM admins WHERE username = $1 AND password = $2`
	err = r.DB.QueryRow(sqlStatement, username, password).Scan(&userId)
	if err != nil {
		return err
	}
	updateSqlStatement := `UPDATE admins SET token = $1 WHERE id = $2`
	_, err = tx.Exec(updateSqlStatement, token, userId)

	if err := tx.Commit(); err != nil {
		return err
	}
	return err
}

// func (r *AdminRepository) Logout(token string) error {
// 	tx, err := r.DB.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	defer func() {
// 		if p := recover(); p != nil {
// 			tx.Rollback()
// 			panic(p)
// 		} else if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	updateSqlStatement := `UPDATE admins SET token = null WHERE token = $1`
// 	_, err = tx.Exec(updateSqlStatement, token)
// 	return err
// }

func (r *AdminRepository) GetByToken(token string) (string, error) {
	sqlStatement := `SELECT token FROM admins WHERE token = $1`
	var userToken string
	err := r.DB.QueryRow(sqlStatement, token).Scan(&userToken)
	if err == sql.ErrNoRows {
		return "", err
	} else if err != nil {
		return "", err
	}
	return userToken, nil
}
