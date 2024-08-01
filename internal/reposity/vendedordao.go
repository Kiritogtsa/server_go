package reposity

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Kiritogtsa/server_go/config"
	"github.com/Kiritogtsa/server_go/internal/domain"
)

type Vendedorinterface interface {
	Persist(*domain.Vendedor) (*domain.Vendedor, error)
	insert(*domain.Vendedor) (*domain.Vendedor, error)
	update(*domain.Vendedor) (*domain.Vendedor, error)
	FindById(int) (*domain.Vendedor, error)
	FindByUserid(int) (*domain.Vendedor, error)
	Delete(int) error
}

type Vendedordao struct {
	Conn config.Config
}

func NewVendedordao(conn config.Config) Vendedorinterface {
	return &Vendedordao{Conn: conn}
}

// Insert insere um novo vendedor no banco de dados
func (coon *Vendedordao) insert(vendedor *domain.Vendedor) (*domain.Vendedor, error) {
	fmt.Println(vendedor)
	dao := coon.Conn.Getdb()
	if dao == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}
	query := "INSERT INTO vendedores (usuario_id, produtos) VALUES (?, ?)"
	stmt, err := dao.Prepare(query)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(vendedor.User, vendedor.Produtos)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Erro ao obter o último ID inserido:", err)
		return nil, err
	}

	vendedor.SetId(int(id))
	return vendedor, nil
}

// Update atualiza um vendedor existente no banco de dados
func (conn *Vendedordao) update(vendedor *domain.Vendedor) (*domain.Vendedor, error) {
	dao := conn.Conn.Getdb()
	if dao == nil {
		return nil, errors.New("nao abriu o db")
	}
	query := "UPDATE vendedores SET usuario_id = ?, produtos = ? WHERE id = ?"
	stmt, err := dao.Prepare(query)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vendedor.ID, vendedor.Produtos, vendedor.ID)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	return vendedor, nil
}

// Persist persiste o vendedor no banco de dados (insere ou atualiza dependendo da existência do ID)
func (dao *Vendedordao) Persist(vendedor *domain.Vendedor) (*domain.Vendedor, error) {
	if vendedor.ID == 0 {
		return dao.insert(vendedor)
	}
	vendedor, err := dao.update(vendedor)
	if err != nil {
		return nil, err
	}
	return vendedor, nil
}

// FindById retorna um vendedor com base no ID
func (coon *Vendedordao) FindById(id int) (*domain.Vendedor, error) {
	dao := coon.Conn.Getdb()
	if dao == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}
	query := "SELECT id, usuario_id, produtos FROM vendedores WHERE id = ?"
	stmt, err := dao.Prepare(query)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	vendedor := &domain.Vendedor{}
	err = stmt.QueryRow(id).Scan(&vendedor.ID, &vendedor.User, &vendedor.Produtos)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	return vendedor, nil
}

func (coon *Vendedordao) FindByUserid(id int) (*domain.Vendedor, error) {
	dao := coon.Conn.Getdb()
	if dao == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}
	query := "SELECT id, usuario_id, produtos FROM vendedores WHERE usuario_id = ?"
	stmt, err := dao.Prepare(query)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	vendedor := &domain.Vendedor{}
	err = stmt.QueryRow(id).Scan(&vendedor.ID, &vendedor.User, &vendedor.Produtos)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Erro ao executar a instrução SQL:", err)
		return nil, err
	}

	return vendedor, nil
}

// Delete remove um vendedor com base no ID
func (coon *Vendedordao) Delete(id int) error {
	dao := coon.Conn.Getdb()
	if dao == nil {
		return nil
	}
	query := "DELETE FROM vendedores WHERE id = ?"
	stmt, err := dao.Prepare(query)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return err
	}

	return nil
}
