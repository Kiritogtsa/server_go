package entries

type User struct {
	ID       int         `json:"id"`
	Name     string      `json:"nome"`
	Email    string      `json:"email"`
	Idade    string      `json:"idade"`
	Password string      `json:"senha"`
	Produtos *[]Produtos `json:"produtos"`
	Vendedor *Vendedor   `json:"vendedor"`
}
