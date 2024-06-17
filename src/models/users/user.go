package users

import (
	"errors"

	"github.com/Kiritogtsa/server_go/src/models/vendedor"
)

type User struct {
	id       int
	name     string
	email    string
	senha    string
	vendedor *vendedor.Vendedor
	saldo    float64
}

func NewUser(nome string, email string, senha string, vendedor *vendedor.Vendedor) (*User, error) {
	if nome == "" || email == "" || senha == "" {
		return nil, errors.New("faltando uma variável (nome, email ou senha)")
	}
	return &User{
		name:     nome,
		email:    email,
		senha:    senha,
		vendedor: vendedor,
	}, nil
}

func (user *User) GetVendedor() *vendedor.Vendedor {
	return user.vendedor
}
func (user *User) SetVendedor(vendedor *vendedor.Vendedor) {
	user.vendedor = vendedor
}

func (user *User) GetSaldo() float64 {
	return user.saldo
}

func (user *User) SetSaldo(saldo float64) {
	user.saldo = saldo
}
func (user *User) SetName(nome string) {
	user.name = nome
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetEmail(email string) {
	user.email = email
}

func (user *User) GetEmail() string {
	return user.email
}

func (user *User) SetSenha(senha string) {
	user.senha = senha
}

func (user *User) GetSenha() string {
	return user.senha
}

func (user *User) SetID(id int) {
	user.id = id
}

func (user *User) GetID() int {
	return user.id
}
