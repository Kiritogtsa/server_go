package repostory

import (
	"database/sql"

	"github.com/Kiritogtsa/server_go/domain/entries"
)

type Usercrud interface {
	save(*entries.User) error
	update(*entries.User) error
	Persti(*entries.User) error
	Delete(int) error
}

type Userdao struct {
	Conn *sql.Conn
}

func Newuserdao(conn *sql.Conn) Usercrud {
	return &Userdao{Conn: conn}
}

func (s *Userdao) save(*entries.User) error {
	return nil
}

func (s *Userdao) update(*entries.User) error {
	return nil
}

func (s *Userdao) Persti(*entries.User) error {
	return nil
}

func (s *Userdao) Delete(int) error {
	return nil
}
