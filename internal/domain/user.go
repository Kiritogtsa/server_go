package domain

import (
	"errors"
)

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Senha      string    `json:"-"`
	Vendedor   *Vendedor `json:"vendedor"`
	Vendedorid int       `json:"vendedorid"`
	Saldo      float64   `json:"saldo"`
	Produtos   []Produtos
}

func NewUser(nome string, email string, senha string, vendedor *Vendedor) (*User, error) {
	if nome == "" || email == "" || senha == "" {
		return nil, errors.New("faltando uma variável (nome, email ou senha)")
	}
	return &User{
		Name:     nome,
		Email:    email,
		Senha:    senha,
		Vendedor: vendedor,
	}, nil
}

func (user *User) GetVendedor() *Vendedor {
	return user.Vendedor
}

func (user *User) SetVendedor(vendedor *Vendedor) {
	user.Vendedor = vendedor
}

func (user *User) GetSaldo() float64 {
	return user.Saldo
}

func (user *User) SetSaldo(saldo float64) {
	user.Saldo = saldo
}

func (user *User) SetName(nome string) {
	user.Name = nome
}

func (user *User) GetName() string {
	return user.Name
}

func (user *User) SetEmail(email string) {
	user.Email = email
}

func (user *User) GetEmail() string {
	return user.Email
}

func (user *User) SetSenha(senha string) {
	user.Senha = senha
}

func (user *User) GetSenha() string {
	return user.Senha
}

func (user *User) SetID(id int) {
	user.ID = id
}

func (user *User) GetID() int {
	return user.ID
}
