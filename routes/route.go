package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/Kiritogtsa/server_go/handles"
)

type Serverfunctions interface {
	Configroutes()
	Runserver()
}

func Serverfiles(w http.ResponseWriter, r *http.Request) {
}

type Server struct {
	Router  chi.Router
	address string
	handle  *handles.Handles
}

func Newserver(port int, address string, handles *handles.Handles) *Server {
	router := chi.NewRouter()
	address = configaddress(port, address)
	return &Server{Router: router, address: address, handle: handles}
}

func (s *Server) routespublic() {
	s.Router.Get("/", s.handle.Getbyalluser)
	s.Router.Post("/user", s.handle.CreateUser)
}

func (s *Server) routesprivates() {
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	s.Router.Put("/user", s.handle.Updateuser)
}

func (s *Server) Configroutes() {
	s.routespublic()
	s.routesprivates()
}

func configaddress(port int, address string) string {
	if port == 0 {
		port = 8080
	}
	if address != "" && address != "localhost" {
		return address + ":" + strconv.Itoa(port)
	}
	return ":" + strconv.Itoa(port)
}

func (s *Server) Runserver() {
	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(s.address, s.Router)
	if err != nil {
		fmt.Println(err)
		return
	}
}
