package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"4_http-server/external"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// Funções handler
func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Você fez um GET")
	case "POST":
		fmt.Fprintf(w, "Você fez um POST")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Método não permitido")
	}
}

func ufsLocationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// acessando APIs externas
	response := external.GetUfsLocationApi()

	// Codificando a resposta em JSON
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao codificar a resposta: %v", err), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Associando as rotas aos handlers
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/api/location/ufs", ufsLocationHandler)

	// Iniciando o servidor na porta 8080
	fmt.Println("Servidor rodando na porta 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
