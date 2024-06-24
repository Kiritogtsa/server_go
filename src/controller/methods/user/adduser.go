package methods

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Kiritogtsa/server_go/src/models/users"
	"github.com/Kiritogtsa/server_go/src/models/vendedor"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println(lines)
	if is_vendedor == "" {
		is_vendedor = "0"
	}
	user, err := users.NewUser(nome, email, senha, nil)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao criar usuário: %v", err), http.StatusBadRequest)
		return
	}
	userdao, err := users.NewUserdao()
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao criar o userdao: %v", err), http.StatusBadRequest)
		return
	}
	user, err = userdao.Persistir(user, is_vendedor)

	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao adicionar o usuario no db: %v", err), http.StatusBadRequest)
		return
	}
	if user.GetID() != 0 {

		if is_vendedor == "1" {
			vendedordao, err := vendedor.NewVendedordao()
			if err != nil {
				http.Error(w, fmt.Sprintf("erro ao criar o userdao: %v", err), http.StatusBadRequest)
				return
			}
			vendedor, err := vendedor.Newvendedor(user.GetID(), 0)
			if err != nil {
				http.Error(w, fmt.Sprintf("erro ao criar o venedor na instacia: %v", err), http.StatusBadRequest)
				return
			}
			vendedor, err = vendedordao.Persist(vendedor)
			if err != nil {
				http.Error(w, fmt.Sprintf("erro ao criar o venedor na instacia: %v", err), http.StatusBadRequest)
				return
			}
			user.SetVendedor(vendedor)
			userdao.Persistir(user, "1")
		}
	} else {
		log.Println("deu algum erro", err)
	}

	fmt.Printf("Usuário criado: %+v\n", user)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso: %s (%s)\n", user.GetName(), user.GetEmail())))
}
