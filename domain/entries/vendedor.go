package entries

type Vendedor struct {
	// primaru key
	ID int `json:"id"`

	// database for lenght of slice
	Produtos []Produtos `json:"Produtos"`

	// get user(Vendedor) for buy to produc
	Userid int `json:"user_id"`
}
