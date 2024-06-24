package produtos

import (
	"errors"

	"github.com/Kiritogtsa/server_go/src/models/users"
)

type Produtos struct {
	ID         int
	Nome       string
	Quantidade int
	vendedor   *users.User
	Preco      float64
}

func NewProduto(Nome string, Quantidade int, vendedor *users.User, preco float64) (*Produtos, error) {
	if Nome != "" || Quantidade != 0 || vendedor != nil {
		return nil, errors.New("nao foi possivel criar um produto")
	}
	return &Produtos{
		Nome:       Nome,
		Quantidade: Quantidade,
		vendedor:   vendedor,
		Preco:      preco,
	}, nil
}
