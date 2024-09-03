package repostory

import (
	"database/sql"

	"github.com/Kiritogtsa/server_go/domain/entries"
)

type Vendedorcrud interface {
	save(*entries.User) error
	update(*entries.User) error
	Persti(*entries.User) error
	Delete(int) error
}

type Vendedordao struct {
	Conn *sql.Conn
}

func Newvendedor(conn *sql.Conn) Vendedorcrud {
	return &Vendedordao{Conn: conn}
}

func (s *Vendedordao) save(*entries.User) error {
	return nil
}

func (s *Vendedordao) update(*entries.User) error {
	return nil
}

func (s *Vendedordao) Persti(*entries.User) error {
	return nil
}

func (s *Vendedordao) Delete(int) error {
	return nil
}
