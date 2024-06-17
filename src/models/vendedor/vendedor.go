package vendedor

import "errors"

type Vendedor struct {
	id       int
	user     int
	produtos int
}

func Newvendedor(user int, produtos int) (*Vendedor, error) {
	if user == 0 {
		return nil, errors.New("faltando uma variável (nome, email ou senha)")
	}
	return &Vendedor{
		user:     user,
		produtos: produtos,
	}, nil
}

func (vendedor *Vendedor) GetId() int {
	return vendedor.id
}
func (vendedor *Vendedor) SetId(id int) {
	vendedor.id = id
}

func (vendedor *Vendedor) GetuserId() int {
	return vendedor.user
}
func (vendedor *Vendedor) SetUserID(id int) {
	vendedor.user = id
}
func (vendedor *Vendedor) Getquant() int {
	return vendedor.produtos
}

func (vendedor *Vendedor) Setquant(quant int) {
	vendedor.produtos = quant
}
