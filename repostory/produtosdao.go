package repostory

import (
	"database/sql"

	"github.com/Kiritogtsa/server_go/domain/entries"
)

type Produtoscrud interface {
	Crud[entries.Produtos]
	GetbyAll() ([]entries.Produtos, error)
}

type Produtosdao struct {
	Conn *sql.Conn
}

func Newprodutosdao(conn *sql.Conn) Produtoscrud {
	return &Produtosdao{Conn: conn}
}

func (s *Produtosdao) save(*entries.Produtos) error {
	return nil
}

func (s *Produtosdao) update(*entries.Produtos) error {
	return nil
}

func (s *Produtosdao) Persti(*entries.Produtos) error {
	return nil
}

func (s *Produtosdao) Delete(int) error {
	return nil
}

// no need to use a instance for vededor, use id_vededor
func (s *Produtosdao) GetbyAll() ([]entries.Produtos, error) {
	return nil, nil
}

// creates more methods for seach
