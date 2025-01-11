package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "adminpassword"
	dbname   = "mydatabase"
)

func ConnDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		user, password, dbname, host, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao verificar conexão com o banco:", err)
		return nil, err
	}
	return db, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(db *sql.DB, username, password string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return fmt.Errorf("erro ao criptografar a senha: %v", err)
	}

	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err = db.Exec(query, username, hash)
	if err != nil {
		return fmt.Errorf("erro ao inserir usuário: %v", err)
	}
	return nil
}
