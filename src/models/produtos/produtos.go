package produtos

import (
	"errors"
	"fmt"

	"github.com/Kiritogtsa/server_go/src/models/users"
)

type Produtos struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	Quantidade int    `json:"quantidade"`
	vendedor   *users.User
	VendedorID int     `json:"vende_id"`
	Preco      float64 `json:"preço"`
}

func NewProduto(
	Nome string,
	Quantidade int,
	vendedor *users.User,
	preco float64,
) (*Produtos, error) {
	fmt.Println(Nome, Quantidade, vendedor, preco)
	if Nome == "" || Quantidade == 0 || vendedor == nil {
		return nil, errors.New("nao foi possivel criar um produto")
	}

	return &Produtos{
		Nome:       Nome,
		Quantidade: Quantidade,
		vendedor:   vendedor,
		Preco:      preco,
	}, nil
}
