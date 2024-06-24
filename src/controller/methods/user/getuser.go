package methods

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kiritogtsa/server_go/src/models/users"
)

func Getall(w http.ResponseWriter, r *http.Request) {
	userdao, err := users.NewUserdao()
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao criar o userdao: %v", err), http.StatusBadRequest)
		return
	}

	// Chama o método SeachAll() para buscar todos os usuários
	userslist, err := userdao.SeachbyALL()
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao buscar todos os usuários: %v", err), http.StatusInternalServerError)
		return
	}
	// Serializa os usuários para JSON
	usersJSON, err := json.Marshal(userslist)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao serializar usuários para JSON: %v", err), http.StatusInternalServerError)
		return
	}
	// Define o cabeçalho Content-Type para JSON
	w.Header().Set("Content-Type", "application/json")
	// Escreve a resposta HTTP com os usuários serializados em JSON
	_, err = w.Write(usersJSON)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao escrever resposta JSON: %v", err), http.StatusInternalServerError)
		return
	}
}
