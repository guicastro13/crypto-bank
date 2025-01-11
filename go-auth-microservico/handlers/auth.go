package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"go-auth-microservico/models"
	"go-auth-microservico/utils"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
		return
	}

	var loginRequest LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Requisição inválida"))
		return
	}

	db, err := utils.ConnDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	user, err := getUserByUsername(db, loginRequest.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Usuário ou senha inválidos"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJwt(loginRequest.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao gerar token"))
		return
	}

	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getUserByUsername(db *sql.DB, username string) (*models.User, error) {
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
