package methods

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/Kiritogtsa/server_go/src/models/users"
)

func GetbyID(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "user_id")
	user_id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da solicitação", http.StatusInternalServerError)
		return
	}
	userdao, err := users.NewUserdao()
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao criar o userdao: %v", err), http.StatusInternalServerError)
		return
	}
	user, err := userdao.Seachbyid(user_id)
	if err != nil {
		http.Error(w, fmt.Sprintf("usuario not found: %v", err), http.StatusNotFound)
		return
	}
	json, err := json.Marshal(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao criar o json: %v", err), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
