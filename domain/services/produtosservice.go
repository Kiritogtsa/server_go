package services

import "github.com/Kiritogtsa/server_go/domain/entries"

type Services_produtos interface {
	services
}

type Servicesprodutos struct {
	produtos *entries.Produtos
}

func Newserviceprodutos(produtos *entries.Produtos) Services_produtos {
	return &Servicesprodutos{produtos: produtos}
}

func (*Servicesprodutos) verificacao() bool {
	return true
}

func (*Servicesprodutos) validacao() bool {
	return true
}

func (s *Servicesprodutos) Create() bool {
	return true
}

func (*Servicesprodutos) Update() bool {
	return true
}

func (*Servicesprodutos) Delete() bool {
	return true
}
