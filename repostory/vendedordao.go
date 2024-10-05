package repostory

import (
	"database/sql"

	"github.com/Kiritogtsa/server_go/domain/entries"
)

type Vendedorcrud interface {
	Crud[entries.Vendedor]
}

type Vendedordao struct {
	Conn *sql.DB
}

func Newvendedor(conn *sql.DB) Vendedorcrud {
	return &Vendedordao{Conn: conn}
}

func (s *Vendedordao) save(*entries.Vendedor) error {
	return nil
}

func (s *Vendedordao) update(*entries.Vendedor) error {
	return nil
}

func (s *Vendedordao) Persti(*entries.Vendedor) error {
	return nil
}

func (s *Vendedordao) Delete(int) error {
	return nil
}
