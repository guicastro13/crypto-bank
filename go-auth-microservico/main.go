package main

import (
	"fmt"
	"net/http"
	"go-auth-microservico/handlers"
	"go-auth-microservico/utils"
)

func main() {
	http.HandleFunc("/health", handlers.HeatlhHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/protected", utils.JwtMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Você acessou uma rota protegida!"))
	}))

	port := "8080"
	fmt.Printf("Servidor rodando na porta %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err !=  nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
		panic(err)
	}
}