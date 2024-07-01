package middlerloggin

import (
	"net/http"

	"github.com/Kiritogtsa/server_go/config"
)

type Middlerlogginin interface {
	Autho(http.Handler) http.Handler
}

type Middlerloggin struct {
}

func NewLogginmiddler() Middlerlogginin {
	return &Middlerloggin{}
}
func (m *Middlerloggin) Autho(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := config.Store.Get(r, "sessao-usuario")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Verifica se o usuário está autenticado
		if session.Values["sessao-usuario"] == nil {
			http.Error(w, "Acesso não autorizado", http.StatusUnauthorized)
			return
		}

		// Continua para o próximo handler se autenticado
		next.ServeHTTP(w, r)
	})
}
