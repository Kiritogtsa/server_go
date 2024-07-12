package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Kiritogtsa/server_go/middlers"
	"github.com/Kiritogtsa/server_go/middlers/middlerloggin"
)

func serveFile(w http.ResponseWriter, r *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Erro ao abrir o arquivo %s: %v", filename, err),
			http.StatusInternalServerError,
		)
		return

	}
	defer file.Close()
	ext := filepath.Ext(filename)
	var contentType string
	switch ext {
	case ".html":
		contentType = "text/html; charset=utf-8"
	case ".htm":
		contentType = "text/html; charset=utf-8"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	default:
		contentType = "application/octet-stream"
	}

	w.Header().Set("Content-Type", contentType)

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Erro ao enviar conteúdo do arquivo %s: %v", filename, err),
			http.StatusInternalServerError,
		)
	}
}

// The function `handleMethod` in Go handles different HTTP methods by serving files for GET requests
// and processing form data for POST requests.
func handleMethod(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}
	filename := "src/view" + path

	serveFile(w, r, filename)
}

func server() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in server function:", r)
			debug.PrintStack()
		}
	}()
	main, err := middlers.Newmainmiddler()
	if err != nil {
		fmt.Println(" deu algum erro no usermiddler", err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// switch routes for mount, in chi at goland
	r.Get("/*", handleMethod)
	r.Group(func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Use(middlerloggin.NewLogginmiddler().Autho)
			user := main.GetUserMiddler()
			user.SetRoutesUser(r)
		})
	})
	r.Group(func(r chi.Router) {
		r.Route("/produto", func(r chi.Router) {
			r.Use(middlerloggin.NewLogginmiddler().Autho)
			main.GetProdutoMiddler().SetRoutesProdutos(r)
		})
	})

	r.Post("/createuser", main.GetUserMiddler().AddUser)
	r.Post("/loggin", main.GetUserMiddler().Loggin)
	fmt.Println("servidor roando em http://localhost:8000")
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	server()
}
