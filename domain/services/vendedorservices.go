package services

import "github.com/Kiritogtsa/server_go/domain/entries"

type Services_vendedor interface {
	Vericacao() bool
	Validacao() bool
}

type Servicesvendedor struct {
	vendedor *entries.Vendedor
}

func Newservicevendedor(Vendedor *entries.Vendedor) Services_vendedor {
	return &Servicesvendedor{vendedor: Vendedor}
}

func (*Servicesvendedor) Vericacao() bool {
	return true
}

func (*Servicesvendedor) Validacao() bool {
	return true
}
