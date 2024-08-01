package domain

import "errors"

type Produtos struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	Quantidade int    `json:"quantidade"`
	VendedorID int     `json:"vende_id"`
	Preco      float64 `json:"preço"`
	Vendedor *User
}
func NewProduto(
	Nome string,
	Quantidade int,
	preco float64,
	vendedor *User,
) (*Produtos, error) {
	if Nome == "" || Quantidade == 0 || vendedor == nil {
		return nil, errors.New("nao foi possivel criar um produto")
	}

	return &Produtos{
		Nome:       Nome,
		Quantidade: Quantidade,
		Vendedor:   vendedor,
		Preco:      preco,
	}, nil
}
