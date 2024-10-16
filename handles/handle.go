package handles

import (
	"database/sql"
	"encoding/json"
	"mime"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	"github.com/Kiritogtsa/server_go/domain/entries"
	"github.com/Kiritogtsa/server_go/repostory"
)

var Store = sessions.NewCookieStore([]byte("seila"))

type Handles struct {
	UserRep *repostory.Usercrud
	ProdRep *repostory.Produtoscrud
}

func Newhandles(c *sql.DB) *Handles {
	userdao := repostory.Newuserdao(c)
	proddao := repostory.Newprodutosdao(c)
	return &Handles{UserRep: &userdao, ProdRep: &proddao}
}

func (Handles) Gettype(r *http.Request, mimetype string) bool {
	contentype := r.Header.Get("Content-type")
	if contentype == "" {
		return mimetype == "application/octet-stream"
	}
	for _, v := range strings.Split(contentype, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}

// nao faz parte da api
// depois ajeitar a forma de receber os dados, para caso vire uma apu
func (h *Handles) CreateUser(w http.ResponseWriter, r *http.Request) {
	ver := h.Gettype(r, "application/json")
	if !ver {
		w.Write([]byte("nao e tipo correto"))
	}
	session, _ := Store.Get(r, "seila")
	var user entries.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("nao foi possivel criar o usuario"))
	}
	userdao := *h.UserRep
	err = userdao.Persti(&user)
	if err != nil {
		w.Write([]byte("nao foi possivel criar o usuario"))
	}
	dados, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte("nao foi possivel criar o usuario"))
	}
	session.Values["user"] = dados
	session.Values["logado"] = true
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home.http", http.StatusSeeOther)
}

// ja fazer os middleware de auteficaçao e permissao e depois testar o codigo antes de continuar o codigo
// nao faz parte da api
func (h *Handles) Login(w http.ResponseWriter, r *http.Request) {
	ver := h.Gettype(r, "application/json")
	if !ver {
		w.Write([]byte("nao e tipo correto"))
	}
	session, _ := Store.Get(r, "seila")
	var user entries.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("json invalido"))
	}
	userdao := *h.UserRep
	u, err := userdao.Getbyname(user.Name)
	if err != nil {
		w.Write([]byte("erro no banco de dados"))
	}
	dados, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte("senha errada"))
	}
	com := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if com != nil {
		http.Error(w, "Senha incorreta", http.StatusUnauthorized)
		return
	}
	session.Values["user"] = dados
	session.Values["logado"] = true
	err = session.Save(r, w)
	if err != nil {
		w.Write([]byte(""))
	}
	http.Redirect(w, r, "/home.http", http.StatusSeeOther)
}

func (h *Handles) Getbyalluser(http.ResponseWriter, *http.Request) {
}
func (j *Handles) Updateuser(http.ResponseWriter, *http.Request) {}
