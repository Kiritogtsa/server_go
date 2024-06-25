package produtos

import (
	"errors"
	"fmt"
	"log"

	"github.com/Kiritogtsa/server_go/config"
	"github.com/Kiritogtsa/server_go/src/models/users"
)

type Produtosinterface interface {
	insert(*Produtos) (*Produtos, error)
	update(*Produtos) (*Produtos, error)
	Persistir(*Produtos) (*Produtos, error)
	Getall() ([]*Produtos, error)
	Getbyid(int) (*Produtos, error)
	Getbyvendedor(*Produtos) (*users.User, error)
}

type ProdutosCrud struct {
	Conn config.Config
}

func NewProdutoCrud(conn config.Config) Produtosinterface {
	return &ProdutosCrud{Conn: conn}
}

func (crud *ProdutosCrud) insert(p *Produtos) (*Produtos, error) {
	db := crud.Conn.Getdb()
	if db == nil {
		return p, errors.New("erro ao obter a conexão com o banco de dados")
	}

	sql := "INSERT INTO produtos (nome, quant, preco, vendedor_id) VALUES (?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return p, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Nome, p.Quantidade, p.Preco, p.vendedor.Vendedorid)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return p, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("erro ao obtem o id de insert", err)
		return nil, err
	}
	p.ID = int(id)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Erro ao obter o número de linhas afetadas:", err)
		return p, err
	}

	fmt.Printf("Inseridos %d registros.\n", rowsAffected)
	return p, nil

}
func (crud *ProdutosCrud) update(*Produtos) (*Produtos, error) {
	return nil, nil
}
func (crud *ProdutosCrud) Persistir(p *Produtos) (*Produtos, error) {
	if p.ID == 0 {
		return crud.insert(p)
	}
	return crud.update(p)
}
func (crud *ProdutosCrud) Getall() ([]*Produtos, error) {
	db := crud.Conn.Getdb()
	if db == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}

	sql := "SELECT id, nome, quant, preco, vendedor_id FROM produtos"
	rows, err := db.Query(sql)
	if err != nil {
		log.Println("Erro ao executar a consulta SQL:", err)
		return nil, err
	}
	defer rows.Close()

	var produtosList []*Produtos
	for rows.Next() {
		var p Produtos
		var vendedorID int
		if err := rows.Scan(&p.ID, &p.Nome, &p.Quantidade, &p.Preco, &vendedorID); err != nil {
			log.Println("Erro ao escanear os resultados:", err)
			return nil, err
		}
		userDAO := users.NewUserdao(crud.Conn)
		vendedor, err := userDAO.GetUserByveid(vendedorID)
		if err != nil {
			log.Println("Erro ao buscar vendedor pelo ID:", err)
			return nil, err
		}
		p.vendedor = vendedor
		produtosList = append(produtosList, &p)
	}

	if err := rows.Err(); err != nil {
		log.Println("Erro ao iterar pelos resultados:", err)
		return nil, err
	}

	return produtosList, nil
}

func (crud *ProdutosCrud) Getbyid(id int) (*Produtos, error) {
	db := crud.Conn.Getdb()
	if db == nil {
		return nil, errors.New("erro ao obter a conexão com o banco de dados")
	}

	sql := "SELECT id, nome, quant, preco, vendedor_id FROM produtos WHERE id = ? "
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return nil, err
	}
	defer stmt.Close()

	var p Produtos

	err = stmt.QueryRow(id).Scan(&p.ID, &p.Nome, &p.Quantidade, &p.Preco, &p.VendedorID)
	if err != nil {
		log.Println("Erro ao executar a quer:", err)
		return nil, err
	}
	userDAO := users.NewUserdao(crud.Conn)
	vendedor, err := userDAO.GetUserByveid(p.VendedorID)
	if err != nil {
		log.Println("Erro ao buscar vendedor pelo ID:", err)
		return nil, err
	}
	p.vendedor = vendedor
	return &p, nil
}
func (crud *ProdutosCrud) Getbyvendedor(*Produtos) (*users.User, error) {
	return nil, nil
}
