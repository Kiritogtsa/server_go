// The main function changes the current working directory to a specified location and then starts an
// HTTP server to serve files and handle POST requests.
package main

// The `import` statement in Go is used to include packages that provide functionality needed in the
// program. In this case, the `import` statement is importing the following packages:
import (
	"fmt"
	"io" // Importa o pacote methods
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"server/src/controller/methods"
	//"server/src/controller/methods"
)

// The serveFile function serves a specified file over HTTP with the appropriate content type based on
// the file extension.
func serveFile(w http.ResponseWriter, r *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao abrir o arquivo %s: %v", filename, err), http.StatusInternalServerError)
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
		http.Error(w, fmt.Sprintf("Erro ao enviar conteúdo do arquivo %s: %v", filename, err), http.StatusInternalServerError)
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

// func handlepostuser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Requisição recebida em /user com método:", r.Method)

// 	// Adicionando log para ler e mostrar o corpo da solicitação
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Erro ao ler o corpo da solicitação", http.StatusInternalServerError)
// 		return
// 	}
// 	defer r.Body.Close()
// 	fmt.Println("Corpo da requisição:", string(body))

// 	// Enviar resposta de volta ao cliente
// 	w.Write([]byte("vem aqui"))
// }

func server() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/*", handleMethod)
	r.Post("/user", methods.AddUser)
	fmt.Println("servidor roando em http://localhost:8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println(err)
	}
}

// The main function retrieves the current working directory, changes the directory to a specified
// location, and then calls the server function.
func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho atual", err, cwd)
		return
	}
	newDir := os.Getenv("HOME") + "/java"
	err = os.Chdir(newDir)
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho atual", err)
		return
	}
	server()
}
