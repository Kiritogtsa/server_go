package services

import (
	"log"

	"github.com/Kiritogtsa/server_go/domain/entries"
	"github.com/Kiritogtsa/server_go/repostory"
)

type Services_user interface {
	verificacao() bool
	validacao() bool
	Createuser() bool
	Updateuser() bool
}

type Servicesuser struct {
	user         *entries.User
	usercrud     *repostory.Usercrud
	produtoscrud *repostory.Produtoscrud
}

func Newserviceuser(
	user *entries.User,
	usercurd *repostory.Usercrud,
	produtoscrud *repostory.Produtoscrud,
) Services_user {
	return &Servicesuser{user: user, usercrud: usercurd, produtoscrud: produtoscrud}
}

func (*Servicesuser) verificacao() bool {
	return true
}

func (*Servicesuser) validacao() bool {
	return true
}

func (s *Servicesuser) Createuser() bool {
	s.verificacao()
	s.validacao()
	usercur := *s.usercrud
	err := usercur.Persti(s.user)
	if err != nil {
		log.Print("nao foi possivel criar o usuario", err)
		return false
	}
	return true
}

func (s *Servicesuser) Updateuser() bool {
	s.verificacao()
	s.validacao()
	usercurd := *s.usercrud
	err := usercurd.Persti(s.user)
	if err != nil {
		log.Print("nao foi possivel criar o usuario", err)
		return false
	}
	return true
}

func (s *Servicesuser) Delete() bool {
	return true
}
