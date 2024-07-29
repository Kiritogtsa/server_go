package middlerprodutos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/Kiritogtsa/server_go/config"
	"github.com/Kiritogtsa/server_go/src/models/produtos"
	"github.com/Kiritogtsa/server_go/src/models/users"
)

type produto struct {
	ID         int
	Nome       string
	Quantidade int
	VendedorID int
	Preco      float64
}
type ProdutosMiddlerinterface interface {
	Insert(http.ResponseWriter, *http.Request)
	Getbyid(http.ResponseWriter, *http.Request)
	Getall(http.ResponseWriter, *http.Request)
	SetRoutesProdutos(chi.Router)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type ProdutosMiddler struct {
	Produtoscrud produtos.Produtosinterface
	conn         config.Config
}

func NewProdutomiddler() (ProdutosMiddlerinterface, error) {
	conn, err := config.NewConn()
	if err != nil {
		return nil, err
	}
	produtoscrud := produtos.NewProdutoCrud(conn)
	return &ProdutosMiddler{
		Produtoscrud: produtoscrud,
		conn:         conn,
	}, nil
}

func (m *ProdutosMiddler) SetRoutesProdutos(r chi.Router) {
	r.Post("/", m.Insert)
	r.Get("/", m.Getall)
	r.Get("/{produto_id}", m.Getbyid)
	r.Post("/{produto_id}", m.Update)
	r.Delete("/{produto_id}", m.Delete)
}

func (m *ProdutosMiddler) Insert(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da solicitação", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var produto produto
	err = json.Unmarshal(body, &produto)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	sessao, err := config.Store.Get(r, "sessao-usuario")
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Erro ao obter a sessao: %v, \n %v", err, sessao),
			http.StatusInternalServerError,
		)
		return
	}
	userData, ok := sessao.Values["sessao-usuario"].([]byte)
	if !ok {
		http.Error(w, "Usuário não encontrado na sessão", http.StatusNotFound)
		return
	}

	var user users.User
	err = json.Unmarshal(userData, &user)
	if err != nil {
		http.Error(w, "Erro ao desserializar usuário", http.StatusInternalServerError)
		return
	}
	fmt.Println(user)
	if user.Vendedor == nil {
		http.Error(w, " usuário nao autorilizado", http.StatusInternalServerError)
		return
	}
	produto2, err := produtos.NewProduto(produto.Nome, produto.Quantidade, &user, produto.Preco)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Erro ao buscar o id do produto: %v, \n %v", err, produto),
			http.StatusInternalServerError,
		)
		return
	}
	produto2, err = m.Produtoscrud.Persistir(produto2)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Erro ao persistir produto: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(produto2)
}

func (m *ProdutosMiddler) Getall(w http.ResponseWriter, r *http.Request) {
	produtoslist, err := m.Produtoscrud.Getall()
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Erro ao buscar todos os produtos: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtoslist)
}

func (m *ProdutosMiddler) Getbyid(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "produto_id")
	produto_id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "ID do produto inválido", http.StatusBadRequest)
		return
	}

	produto, err := m.Produtoscrud.Getbyid(produto_id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Produto não encontrado: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produto)
}

func (m *ProdutosMiddler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("vem aqui")
	pr := chi.URLParam(r, "produto_id")
	pi, err := strconv.Atoi(pr)
	if err != nil {
		http.Error(w, "erro ao obter o id", http.StatusBadGateway)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da solicitação", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var produto produto
	err = json.Unmarshal(body, &produto)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da solicitação", http.StatusInternalServerError)
		return
	}
	sessao, err := config.Store.Get(r, "sessao-usuario")
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Erro ao obter a sessao: %v, \n %v", err, sessao),
			http.StatusInternalServerError,
		)
		return
	}
	userData, ok := sessao.Values["sessao-usuario"].([]byte)
	if !ok {
		http.Error(w, "Usuário não encontrado na sessão", http.StatusNotFound)
		return
	}

	var user users.User
	err = json.Unmarshal(userData, &user)
	if err != nil {
		http.Error(w, "Erro ao desserializar usuário", http.StatusInternalServerError)
		return
	}
	p := produtos.Produtos{
		ID:         pi,
		Nome:       produto.Nome,
		Quantidade: produto.Quantidade,
		Preco:      produto.Preco,
	}
	produtoc, err := m.Produtoscrud.Getbyid(p.ID)
	if err != nil {
		http.Error(w, "erro ao buscar o produto original", http.StatusInternalServerError)
		return
	}
	if produtoc.VendedorID != user.Vendedorid {
		http.Error(w, "nao eo vendedor do produto", http.StatusInternalServerError)
		return
	}
	pd, err := m.Produtoscrud.Persistir(&p)
	if err != nil {
		http.Error(w, "erro ao persirtir o objeto", http.StatusInternalServerError)
		return
	}
	fmt.Println("sucesso: ", pd)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (m *ProdutosMiddler) Delete(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "produto_id")
	pri, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "erro ao pegar o id", http.StatusBadRequest)
		return
	}
	pd, err := m.Produtoscrud.Getbyid(pri)
	if err != nil {
		http.Error(w, "erro ao pegar o usuario no bd", http.StatusBadRequest)
		return
	}
	err = m.Produtoscrud.DeleteProduto(pd)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "erro ao deletar o produto", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pd)
}
