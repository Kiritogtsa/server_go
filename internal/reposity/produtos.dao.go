package reposity

import (
	"errors"
	"fmt"
	"log"

	"github.com/Kiritogtsa/server_go/config"
	"github.com/Kiritogtsa/server_go/internal/domain"
)

type Produtosinterface interface {
	insert(*domain.Produtos) (*domain.Produtos, error)
	update(*domain.Produtos) (*domain.Produtos, error)
	Persistir(*domain.Produtos) (*domain.Produtos, error)
	Getall() ([]*domain.Produtos, error)
	Getbyid(int) (*domain.Produtos, error)
	Getbyvendedor(*domain.Produtos) (*domain.User, error)
	DeleteProduto(*domain.Produtos) error
}

type ProdutosCrud struct {
	Conn config.Config
}

func NewProdutoCrud(conn config.Config) Produtosinterface {
	return &ProdutosCrud{Conn: conn}
}

func (crud *ProdutosCrud) insert(p *domain.Produtos) (*domain.Produtos, error) {
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

	result, err := stmt.Exec(p.Nome, p.Quantidade, p.Preco, p.Vendedor.Vendedorid)
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

func (crud *ProdutosCrud) update(p *domain.Produtos) (*domain.Produtos, error) {
	db := crud.Conn.Getdb()
	if db == nil {
		return p, errors.New("erro ao obter a conexão com o banco de dados")
	}

	sql := "update produtos set nome = ?, quant=?, preco=? where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Erro ao preparar a instrução SQL:", err)
		return p, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Nome, p.Quantidade, p.Preco, p.ID)
	if err != nil {
		log.Println("Erro ao executar a instrução SQL:", err)
		return p, err
	}
	log.Println("update succes")
	return p, nil
}

func (crud *ProdutosCrud) Persistir(p *domain.Produtos) (*domain.Produtos, error) {
	if p.ID == 0 {
		return crud.insert(p)
	}
	return crud.update(p)
}

func (crud *ProdutosCrud) Getall() ([]*domain.Produtos, error) {
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

	var produtosList []*domain.Produtos
	for rows.Next() {
		var p domain.Produtos
		var vendedorID int
		if err := rows.Scan(&p.ID, &p.Nome, &p.Quantidade, &p.Preco, &vendedorID); err != nil {
			log.Println("Erro ao escanear os resultados:", err)
			return nil, err
		}
		userDAO := NewUserdao(crud.Conn)
		vendedor, err := userDAO.GetUserByveid(vendedorID)
		if err != nil {
			log.Println("Erro ao buscar vendedor pelo ID:", err)
			return nil, err
		}
		p.Vendedor = vendedor
		p.VendedorID = p.Vendedor.Vendedorid
		produtosList = append(produtosList, &p)
	}

	if err := rows.Err(); err != nil {
		log.Println("Erro ao iterar pelos resultados:", err)
		return nil, err
	}

	return produtosList, nil
}

func (crud *ProdutosCrud) Getbyid(id int) (*domain.Produtos, error) {
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

	var p domain.Produtos

	err = stmt.QueryRow(id).Scan(&p.ID, &p.Nome, &p.Quantidade, &p.Preco, &p.VendedorID)
	if err != nil {
		log.Println("Erro ao executar a quer:", err)
		return nil, err
	}
	userDAO := NewUserdao(crud.Conn)
	vendedor, err := userDAO.GetUserByveid(p.VendedorID)
	if err != nil {
		log.Println("Erro ao buscar vendedor pelo ID:", err)
		return nil, err
	}
	p.Vendedor = vendedor
	p.VendedorID = p.Vendedor.Vendedorid
	return &p, nil
}

func (crud *ProdutosCrud) Getbyvendedor(*domain.Produtos) (*domain.User, error) {
	return nil, nil
}

func (crud *ProdutosCrud) DeleteProduto(p *domain.Produtos) error {
	if p.ID != 0 {
		db := crud.Conn.Getdb()
		if db == nil {
			return errors.New("nao foi ṕossivel abrir a conexao com o db")
		}
		sql := "delete from produtos where id = ?"
		stmt, err := db.Prepare(sql)
		if err != nil {
			return errors.New("erro ao preparar o sql")
		}
		result, err := stmt.Query(p.ID)
		if err != nil {
			return errors.New("erro ao executar o sql")
		}
		fmt.Println(result)
		return nil
	} else {
		return errors.New("nao foi possivel obter o id do produto")
	}
}
