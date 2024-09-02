package services

import "github.com/Kiritogtsa/server_go/domain/entries"

type Services_user interface {
	Vericacao() bool
	Validacao() bool
}

type Servicesuser struct {
	user *entries.User
}

func Newserviceuser(user *entries.User) Services_user {
	return &Servicesuser{user: user}
}

func (*Servicesuser) Vericacao() bool {
	return true
}

func (*Servicesuser) Validacao() bool {
	return true
}
