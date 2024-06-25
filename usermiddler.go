package middleruser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server_go/src/models"
	"server_go/src/models/users"
	"server_go/src/models/vendedor"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

const Dns string = "root:1234@tcp(127.0.0.1:3306)/loja"

type Usermiddlerinterface interface {
	AddUser(http.ResponseWriter, *http.Request)
	Getall(http.ResponseWriter, *http.Request)
	Getbyid(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
}
type Usermiddler struct {
	Userdao     users.Userdaointerface
	Vendedordao vendedor.Vendedorinterface
}

func NewUserMiddler() (Usermiddlerinterface, error) {
	conn, err := models.NewConn(Dns)
	if err != nil {
		return nil, err
	}
	userdao := users.NewUserdao(conn)
	vendedor := vendedor.NewVendedordao(conn)
	return &Usermiddler{
		Userdao:     userdao,
		Vendedordao: vendedor,
	}, nil
}
func (m *Usermiddler) AddUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da solicitação", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	bodyStr := string(body)
	lines := strings.Split(bodyStr, "&")

	var nome, email, senha, is_vendedor string
	//var is_vendedor bool
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) < 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		switch key {
		case "name":
			nome = value
		case "email":
			email = value
		case "senha":
			senha = value
		case "is_vendedor":
			is_vendedor = value
		}
	}
	user, err := users.NewUser(nome, email, senha, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao criar usuário: %v", err), http.StatusBadRequest)
		return
	}
	user, err = m.Userdao.Persistir(user, "")
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao persitir usuário: %v", err), http.StatusBadRequest)
		return
	}
	if is_vendedor != "" {
		vendedor, err := vendedor.Newvendedor(user.ID, 0)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao criar usuário: %v", err), http.StatusBadRequest)
			return
		}
		user.Vendedor = vendedor
		user, err = m.Userdao.Persistir(user, "")
		if err != nil {
			http.Error(w, fmt.Sprintf("erro ao setar o id do vendedor: %v", err), http.StatusBadRequest)
			return
		}
	}
	fmt.Printf("Usuário criado: %+v\n", user)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso: %s (%s)\n", user.GetName(), user.GetEmail())))
}

func (m *Usermiddler) Getall(w http.ResponseWriter, r *http.Request) {

	// Chama o método SeachAll() para buscar todos os usuários
	userslist, err := m.Userdao.SeachbyALL()
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
func (m *Usermiddler) Getbyid(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "user_id")
	user_id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da solicitação", http.StatusInternalServerError)
		return
	}

	user, err := m.Userdao.Seachbyid(user_id)
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
func (m *Usermiddler) Update(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content-Type deve ser application/json", http.StatusUnsupportedMediaType)
		return
	}
	var user users.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	user_update, err := m.Userdao.Persistir(user)
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
