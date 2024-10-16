package repostory

import (
	"context"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/Kiritogtsa/server_go/domain/entries"
)

type Usercrud interface {
	Crud[entries.User]
	Getbyname(string) (*entries.User, error)
	Getbyemail(string) (*entries.User, error)
}

type Userdao struct {
	Conn *sql.DB

	verdedor Vendedorcrud
}

func Newuserdao(conn *sql.DB) Usercrud {
	return &Userdao{Conn: conn, verdedor: Newvendedor(conn)}
}

// create a user and Vendedor
// Vendedor only create if user as vendedor
func (s *Userdao) save(user *entries.User) error {
	sql := "insert into () values(?,?)"
	ctx := context.Background()
	Password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 2)
	if err != nil {
		log.Print("nao foi possivel iniciar a transaçao com o banco de dados", err)
	}
	tx, err := s.Conn.BeginTx(ctx, nil)
	if err != nil {
		log.Print("nao foi possivel iniciar a transaçao com o banco de dados", err)
	}
	_, err = tx.ExecContext(ctx, sql, user.Name, user.Email, Password)
	if err != nil {
		log.Print("nao foi possivel commitar a transaçao com o banco de dados", err)
	}
	if user.Vendedor != nil {
		err = s.verdedor.Persti(user.Vendedor)
		if err != nil {
			log.Print("nao foi possivel commitar a transaçao com o banco de dados", err)
		}
		err = s.Persti(user)
		if err != nil {
			log.Print("nao foi possivel commitar a transaçao com o banco de dados", err)
		}
	}
	return nil
}

// it connot to udpate for produtos
// he can to update vendedor
func (s *Userdao) update(*entries.User) error {
	return nil
}

// check the id, if exist id update he user if not exist create he user
func (s *Userdao) Persti(u *entries.User) error {
	if u.ID == 0 {
		return s.save(u)
	}
	return s.update(u)
}

// the function can delete auser
func (s *Userdao) Delete(int) error {
	return nil
}

// get a user of name, need to get tje produtos of user and get products he have to sell
func (s *Userdao) Getbyname(name string) (*entries.User, error) {
	return nil, nil
}

// get a user of email, need to get tje produtos of user and get products he have to sell
func (s *Userdao) Getbyemail(email string) (*entries.User, error) {
	return nil, nil
}
