package services

import "github.com/Kiritogtsa/server_go/domain/entries"

type Services_vendedor interface {
	services
}

type Servicesvendedor struct {
	vendedor *entries.Vendedor
}

func Newservicevendedor(Vendedor *entries.Vendedor) Services_vendedor {
	return &Servicesvendedor{vendedor: Vendedor}
}

func (*Servicesvendedor) verificacao() bool {
	return true
}

func (*Servicesvendedor) validacao() bool {
	return true
}

func (s *Servicesvendedor) Create() bool {
	return true
}

func (*Servicesvendedor) Update() bool {
	return true
}

func (*Servicesvendedor) Delete() bool {
	return true
}
