package handles

import (
	"database/sql"
	"encoding/json"
	"mime"
	"net/http"
	"strings"

	"github.com/Kiritogtsa/server_go/domain/entries"
	"github.com/Kiritogtsa/server_go/repostory"
)

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

// depois ajeitar a forma de receber os dados, para caso vire uma apu
func (h *Handles) CreateUser(w http.ResponseWriter, r *http.Request) {
	ver := h.Gettype(r, "application/json")
	if !ver {
		w.Write([]byte("nao e tipo correto"))
	}
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
}

func (h *Handles) Getbyalluser(http.ResponseWriter, *http.Request) {
}
func (j *Handles) Updateuser(http.ResponseWriter, *http.Request) {}
