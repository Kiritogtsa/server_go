package users

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"server/src/models"
	"server/src/models/vendedor"
)

const Dns string = "root:1234@tcp(127.0.0.1:3306)/loja"

// Userdao é a estrutura para lidar com operações do usuário no banco de dados.
type Userdao struct {
	Conn *models.Conn
}

// NewUserdao cria uma nova instância de Userdao.
func NewUserdao() (*Userdao, error) {
	conn, err := models.NewConn(Dns)
	if err != nil {
		return nil, errors.New("erro ao criar a conexão com o banco de dados")
	}
	return &Userdao{Conn: conn}, nil
}

// insert insere um novo usuário no banco de dados.
func (userdao *Userdao) insert(user *User, is_vendedor string) (*Userdao, error) {
	db := userdao.Conn.Getdb()
	if db == nil {
		return userdao, errors.New("erro ao obter a conexão com o banco de dados")
	}

	hashsenha, err := bcrypt.GenerateFromPassword([]byte(user.GetSenha()), bcrypt.DefaultCost)
	if err != nil {
		return userdao, err
	}
	user.SetSenha(string(hashsenha))

	sql := "INSERT INTO usuario (nome, email, senha, vendedor, vendedor_id, saldo) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return userdao, err
	}
	defer stmt.Close()

	var vendedorID interface{} = nil
	vendedor := user.GetVendedor()
	if vendedor != nil {
		vendedorID = vendedor.GetId()
	}

	result, err := stmt.Exec(user.GetName(), user.GetEmail(), user.GetSenha(), is_vendedor, vendedorID, user.GetSaldo())
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return userdao, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Erro ao obter o número de linhas afetadas:", err)
		return userdao, err
	}

	fmt.Printf("Inseridos %d registros.\n", rowsAffected)
	return userdao, nil
}

// update atualiza um usuário existente no banco de dados.
func (userdao *Userdao) update(user *User, is_vendedor string) (*Userdao, error) {
	db := userdao.Conn.Getdb()
	if db == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}

	hashsenha, err := bcrypt.GenerateFromPassword([]byte(user.GetSenha()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.SetSenha(string(hashsenha))

	// Determinar o `vendedor_id`
	var vendedorID interface{} = nil
	vendedor := user.GetVendedor()
	if vendedor != nil {
		vendedorID = vendedor.GetId()
	}
	fmt.Println("aqui update usuarui")

	// Atualização com possibilidade de NULL para `vendedor_id`
	sql := "UPDATE usuario SET nome = ?, email = ?, senha = ?,vendedor = ?, vendedor_id = ?, saldo = ? WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	fmt.Println("aqui update depois")

	_, err = stmt.Exec(user.GetName(), user.GetEmail(), user.GetSenha(), is_vendedor, vendedorID, user.GetSaldo(), user.GetID())
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	fmt.Printf("Registro com ID %d atualizado com sucesso.\n", user.GetID())
	return userdao, nil
}

// Persistir insere ou atualiza o usuário no banco de dados.
func (userdao *Userdao) Persistir(user *User, is_vendedor string) (*Userdao, error) {
	if user.GetID() == 0 {
		return userdao.insert(user, is_vendedor)
	}
	return userdao.update(user, is_vendedor)
}

// Seachbyid busca um usuário pelo ID no banco de dados.
func (userdao *Userdao) Seachbyid(id int) (*User, error) {
	db := userdao.Conn.Getdb()
	if db == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}

	sql := "SELECT id, nome, email, senha, vendedor_id, saldo FROM usuario WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	user := &User{}
	var vendedorID *int

	err = stmt.QueryRow(id).Scan(&user.id, &user.name, &user.email, &user.senha, &vendedorID, &user.saldo)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}
	if vendedorID != nil { // Verifica se o valor não é nulo
		Vendedordao, err := vendedor.NewVendedordao()
		if err != nil {
			return user, err
		}
		vendedor, err := Vendedordao.FindById(*vendedorID)
		if err != nil {
			return user, err
		}
		user.SetVendedor(vendedor)
	}
	return user, nil
}
func (userdao *Userdao) SeachbyName(name string) (*User, error) {
	db := userdao.Conn.Getdb()
	if db == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}

	sql := "SELECT id, nome, email, senha, vendedor_id, saldo FROM usuario WHERE nome = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	user := &User{}

	var vendedorID *int

	err = stmt.QueryRow(name).Scan(&user.id, &user.name, &user.email, &user.senha, &vendedorID, &user.saldo)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	if vendedorID != nil { // Verifica se o valor não é nulo
		Vendedordao, err := vendedor.NewVendedordao()
		if err != nil {
			return user, err
		}
		vendedor, err := Vendedordao.FindById(*vendedorID)
		if err != nil {
			return user, err
		}
		user.SetVendedor(vendedor)
	}
	fmt.Println("termina aqui")
	return user, nil
}
