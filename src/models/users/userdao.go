package users

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/Kiritogtsa/server_go/config"
	"github.com/Kiritogtsa/server_go/src/models/vendedor"
)

const Dns string = "root:1234@tcp(127.0.0.1:3306)/loja"

type Userdaointerface interface {
	insert(*User, string) (*User, error)
	update(*User) (*User, error)
	Persistir(*User, string) (*User, error)
	Seachbyid(int) (*User, error)
	SeachbyName(string) (*User, error)
	SeachbyALL() ([]*User, error)
	GetUserByveid(int) (*User, error)
	GetUserbyname(string) (*User, error)
	GetUserbyemail(string) (*User, error)
}
type Userdao struct {
	Conn config.Config
}

// NewUserdao cria uma nova instância de Userdao.
func NewUserdao(conn config.Config) Userdaointerface {
	return &Userdao{Conn: conn}
}

// insert insere um novo usuário no banco de dados.
func (userdao *Userdao) insert(user *User, is_vendedor string) (*User, error) {
	var teste int
	if is_vendedor == "" {
		teste = 0
	} else {
		teste = 1
	}
	fmt.Println(is_vendedor)
	db := userdao.Conn.Getdb()
	if db == nil {
		return user, errors.New("erro ao obter a conexão com o banco de dados")
	}

	hashsenha, err := bcrypt.GenerateFromPassword([]byte(user.GetSenha()), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.SetSenha(string(hashsenha))

	sql := "INSERT INTO usuario (nome, email, senha, vendedor,vendedor_id, saldo) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return user, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.GetName(), user.GetEmail(), user.GetSenha(), teste, nil, user.GetSaldo())
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return user, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("erro ao obtem o id de insert", err)
		return nil, err
	}
	user.SetID(int(id))
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Erro ao obter o número de linhas afetadas:", err)
		return user, err
	}

	fmt.Printf("Inseridos %d registros.\n", rowsAffected)
	return user, nil
}

// update atualiza um usuário existente no banco de dados.
func (userdao *Userdao) update(user *User) (*User, error) {
	db := userdao.Conn.Getdb()
	if db == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}

	// Determinar o `vendedor_id`
	var vendedorID interface{} = nil
	vendedor := user.GetVendedor()
	if vendedor != nil {
		vendedorID = vendedor.GetId()
	}
	fmt.Println("aqui update usuarui")

	// Atualização com possibilidade de NULL para `vendedor_id`
	sql := "UPDATE usuario SET nome = ?, email = ?, senha = ?, vendedor_id = ?, saldo = ? WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	fmt.Println("aqui update depois")

	_, err = stmt.Exec(user.GetName(), user.GetEmail(), user.GetSenha(), vendedorID, user.GetSaldo(), user.GetID())
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	fmt.Printf("Registro com ID %d atualizado com sucesso.\n", user.GetID())
	return user, nil
}

// Persistir insere ou atualiza o usuário no banco de dados.
func (userdao *Userdao) Persistir(user *User, is_vendedor string) (*User, error) {
	if user.GetID() == 0 {
		return userdao.insert(user, is_vendedor)
	}
	return userdao.update(user)
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

	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Senha, &vendedorID, &user.Saldo)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}
	if vendedorID != nil { // Verifica se o valor não é nulo
		Vendedordao := vendedor.NewVendedordao(userdao.Conn)
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

	err = stmt.QueryRow(name).Scan(&user.ID, &user.Name, &user.Email, &user.Senha, &vendedorID, &user.Saldo)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	if vendedorID != nil { // Verifica se o valor não é nulo
		Vendedordao := vendedor.NewVendedordao(userdao.Conn)

		vendedor, err := Vendedordao.FindById(*vendedorID)
		if err != nil {
			return user, err
		}
		user.SetVendedor(vendedor)
	}
	fmt.Println("termina aqui")
	return user, nil
}

// SeachAll busca todos os usuários no banco de dados.
func (userdao *Userdao) SeachbyALL() ([]*User, error) {
	db := userdao.Conn.Getdb()
	if db == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}

	sql := "SELECT id, nome, email, senha, vendedor_id, saldo FROM usuario"
	rows, err := db.Query(sql)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}
	defer rows.Close()

	users := []*User{}
	for rows.Next() {
		user := &User{}
		var vendedorID *int
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Senha, &vendedorID, &user.Saldo)
		if err != nil {
			log.Println("Erro ao escanear linha:", err)
			return nil, err
		}

		if vendedorID != nil {
			Vendedordao := vendedor.NewVendedordao(userdao.Conn)

			vendedor, err := Vendedordao.FindById(*vendedorID)
			if err != nil {
				return users, err
			}
			user.SetVendedor(vendedor)
			user.Vendedorid = *vendedorID
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Erro ao iterar sobre os resultados:", err)
		return nil, err
	}

	return users, nil
}

func (crud *Userdao) GetUserByveid(id int) (*User, error) {
	fmt.Println("chega aqui")
	db := crud.Conn.Getdb()
	if db == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}

	sql := "SELECT id, nome, email, senha, vendedor_id, saldo FROM usuario WHERE vendedor_id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	user := &User{}

	var vendedorID *int

	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Senha, &vendedorID, &user.Saldo)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	if vendedorID != nil { // Verifica se o valor não é nulo
		Vendedordao := vendedor.NewVendedordao(crud.Conn)

		vendedor, err := Vendedordao.FindById(*vendedorID)
		if err != nil {
			return user, err
		}
		user.SetVendedor(vendedor)
		user.Vendedorid = user.Vendedor.GetId()
	}
	fmt.Println(user)
	fmt.Println("termina aqui")
	return user, nil
}

func (crud *Userdao) GetUserbyemail(email string) (*User, error) {
	db := crud.Conn.Getdb()
	if db == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}

	sql := "SELECT id, nome, email, senha, vendedor_id, saldo FROM usuario WHERE email = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	user := &User{}
	var vendedorID *int

	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Senha, &vendedorID, &user.Saldo)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	if vendedorID != nil { // Verifica se o valor não é nulo
		Vendedordao := vendedor.NewVendedordao(crud.Conn)

		vendedor, err := Vendedordao.FindById(*vendedorID)
		if err != nil {
			return user, err
		}
		user.SetVendedor(vendedor)
		user.Vendedorid = user.Vendedor.GetId()
	}

	fmt.Println("termina aqui")
	return user, nil
}
func (crud *Userdao) GetUserbyname(name string) (*User, error) {
	db := crud.Conn.Getdb()
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

	err = stmt.QueryRow(name).Scan(&user.ID, &user.Name, &user.Email, &user.Senha, &vendedorID, &user.Saldo)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	if vendedorID != nil { // Verifica se o valor não é nulo
		Vendedordao := vendedor.NewVendedordao(crud.Conn)

		vendedor, err := Vendedordao.FindById(*vendedorID)
		if err != nil {
			return user, err
		}
		user.SetVendedor(vendedor)
		user.Vendedorid = user.Vendedor.GetId()
	}

	fmt.Println("termina aqui")
	return user, nil
}
