package middleruser

import (
	"fmt"
	"io"
	"net/http"
	"server_go/src/models"
	"server_go/src/models/users"
	"server_go/src/models/vendedor"
	"strings"
)

const Dns string = "root:1234@tcp(127.0.0.1:3306)/loja"

type Usermiddlerinterface interface {
	AddUser(http.ResponseWriter, *http.Request)
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
