package domain

import "errors"

type Vendedor struct {
	ID       int
	User     int
	Produtos int
}

func Newvendedor(user int, produtos int) (*Vendedor, error) {
	if user == 0 {
		return nil, errors.New("faltando uma variável (nome, email ou senha)")
	}
	return &Vendedor{
		User:     user,
		Produtos: produtos,
	}, nil
}

func (vendedor *Vendedor) SetId(id int) {
	vendedor.ID = id
}
