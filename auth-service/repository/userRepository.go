package repository

import (
	"auth-service/models"
	"database/sql"
	"fmt"
)

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	query := `SELECT id, username, password FROM users WHERE username = $1`
	row := db.QueryRow(query, username)
	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuário não encontrado")
		}
		return nil, err
	}
	return &user, nil
}
