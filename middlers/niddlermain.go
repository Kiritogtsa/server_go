package middlers

import (
	"github.com/Kiritogtsa/server_go/middlers/middlerprodutos"
	"github.com/Kiritogtsa/server_go/middlers/middliruser"
)

type mainmiddler interface {
	GetUserMiddler() *middliruser.Usermiddlerinterface
	GetProdutoMiddler() *middlerprodutos.ProdutosMiddlerinterface
}

type Main struct {
	usermiddler    middliruser.Usermiddlerinterface
	produtomiddler middlerprodutos.ProdutosMiddlerinterface
}

func New() (mainmiddler, error) {
	usermiddler, err := middliruser.NewUserMiddler()
	if err != nil {
		return nil, err
	}
	produtomiddler, err := middlerprodutos.NewProdutomiddler()
	if err != nil {
		return nil, err
	}
	return &Main{
		usermiddler:    usermiddler,
		produtomiddler: produtomiddler,
	}, nil
}
func (m *Main) GetProdutoMiddler() *middlerprodutos.ProdutosMiddlerinterface {
	return &m.produtomiddler
}
func (m *Main) GetUserMiddler() *middliruser.Usermiddlerinterface {
	return &m.usermiddler
}
