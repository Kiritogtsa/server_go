package services

import "github.com/Kiritogtsa/server_go/domain/entries"

type Services_produtos interface {
	Vericacao() bool
	Validacao() bool
}

type Servicesprodutos struct {
	produtos *entries.Produtos
}

func Newserviceprodutos(produtos *entries.Produtos) Services_produtos {
	return &Servicesprodutos{produtos: produtos}
}

func (*Servicesprodutos) Vericacao() bool {
	return true
}

func (*Servicesprodutos) Validacao() bool {
	return true
}
