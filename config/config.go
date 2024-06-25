package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const Dns string = "root:1234@tcp(127.0.0.1:3306)/loja"

type Config interface {
	Getdb() *sql.DB
}

type Conn struct {
	DB *sql.DB
}

func (db Conn) Getdb() *sql.DB {
	return db.DB
}

func NewConn() (Config, error) {
	// Cria uma conexão com o banco de dados.
	if Dns == "" {
		return nil, fmt.Errorf("DSN não pode estar vazio")
	}
	db, err := sql.Open("mysql", Dns)
	if err != nil {
		return nil, fmt.Errorf("não foi possível se conectar ao banco de dados: %v", err)
	}

	// Verifica se a conexão está ativa.
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("não foi possível verificar a conexão com o banco de dados: %v", err)
	}

	return &Conn{DB: db}, nil
}
