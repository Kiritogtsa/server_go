package middlers

import (
	"github.com/Kiritogtsa/server_go/middlers/middlerprodutos"
	"github.com/Kiritogtsa/server_go/middlers/middliruser"
)

type Mainmiddler interface {
	GetUserMiddler() middliruser.Usermiddlerinterface
	GetProdutoMiddler() middlerprodutos.ProdutosMiddlerinterface
}

type Main struct {
	Usermiddler    middliruser.Usermiddlerinterface
	Produtomiddler middlerprodutos.ProdutosMiddlerinterface
}

func Newmainmiddler() (Mainmiddler, error) {
	usermiddler, err := middliruser.NewUserMiddler()
	if err != nil {
		return nil, err
	}
	produtomiddler, err := middlerprodutos.NewProdutomiddler()
	if err != nil {
		return nil, err
	}
	return &Main{
		Usermiddler:    usermiddler,
		Produtomiddler: produtomiddler,
	}, nil
}
func (m *Main) GetProdutoMiddler() middlerprodutos.ProdutosMiddlerinterface {
	return m.Produtomiddler
}
func (m *Main) GetUserMiddler() middliruser.Usermiddlerinterface {
	return m.Usermiddler
}
