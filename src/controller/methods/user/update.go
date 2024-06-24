package methods

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kiritogtsa/server_go/src/models/users"
)

func Update(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content-Type deve ser application/json", http.StatusUnsupportedMediaType)
		return
	}
	var user users.User
	userdao, err := users.NewUserdao()
	if err != nil {
		http.Error(w, "Erro ao criar a conexao com o bd", http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	user_update, err := userdao.Persistir(&user, "")
	if err != nil {
		http.Error(w, "Erro ao dar update", http.StatusBadRequest)
		return
	}
	json, err := json.Marshal(user_update)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(json)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao escrever resposta JSON: %v", err), http.StatusInternalServerError)
		return
	}
}
