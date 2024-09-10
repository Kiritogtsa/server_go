package repostory

import (
	"database/sql"

	"github.com/Kiritogtsa/server_go/domain/entries"
)

type Usercrud interface {
	Crud[entries.User]
	Getbyname(string) (*entries.User, error)
	Getbyemail(string) (*entries.User, error)
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

func (s *Userdao) Getbyname(name string) (*entries.User, error) {
	return nil, nil
}

func (s *Userdao) Getbyemail(email string) (*entries.User, error) {
	return nil, nil
}
