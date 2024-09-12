package entries

type Produtos struct {
	ID         int  `json:"produto_id"`
	Quantidade int  `json:"quantidade"`
	Vendedor   User `json:"vendedor"`
	Preco      float64
}
