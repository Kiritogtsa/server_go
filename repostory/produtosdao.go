package repostory

import (
	"database/sql"

	"github.com/Kiritogtsa/server_go/domain/entries"
)

type Produtoscrud interface {
	save(*entries.User) error
	update(*entries.User) error
	Persti(*entries.User) error
	Delete(int) error
}

type Produtosdao struct {
	Conn *sql.Conn
}

func Newprodutosdao(conn *sql.Conn) Produtoscrud {
	return &Produtosdao{Conn: conn}
}

func (s *Produtosdao) save(*entries.User) error {
	return nil
}

func (s *Produtosdao) update(*entries.User) error {
	return nil
}

func (s *Produtosdao) Persti(*entries.User) error {
	return nil
}

func (s *Produtosdao) Delete(int) error {
	return nil
}
