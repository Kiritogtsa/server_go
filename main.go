package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Kiritogtsa/server_go/handles"
	"github.com/Kiritogtsa/server_go/routes"
)

type sqldados struct {
	port     string
	host     string
	user     string
	password string
	dbname   string
}

func newdados(port int, host string, user string, password string, dbname string) sqldados {
	return sqldados{port: string(port), host: host, user: user, password: password, dbname: dbname}
}

func (s sqldados) GetConnectionURL() string {
	return fmt.Sprintf(
		"mysql://%s:%s@%s:%s/%s?sslmode=disable",
		s.user,
		s.password,
		s.host,
		s.port,
		s.dbname,
	)
}

func connection(url string) *sql.DB {
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	return db
}

func main() {
	sqldados := newdados(3306, "localhost", "root", "", "loja")
	url := sqldados.GetConnectionURL()
	conn := connection(url)
	handels := handles.Newhandles(conn)
	var server routes.Serverfunctions = routes.Newserver(8080, "", handels)
	server.Configroutes()
	server.Runserver()
}
