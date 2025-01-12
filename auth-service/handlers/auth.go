package handlers

import (
	"encoding/json"
	"net/http"

	"auth-service/database"
	"auth-service/repository"
	"auth-service/utils"

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

	db, err := database.NewDBConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	user, err := repository.GetUserByUsername(db, loginRequest.Username)
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
