package models

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Conn representa a estrutura para a conexão com o banco de dados.
type Conn struct {
	DB *sql.DB
}

func (db Conn) Getdb() *sql.DB {
	return db.DB
}

// NewConn cria uma nova conexão com o banco de dados e a retorna.
func NewConn(dsn string) (*Conn, error) {
	// Cria uma conexão com o banco de dados.
	if dsn == "" {
		return nil, fmt.Errorf("DSN não pode estar vazio")
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("não foi possível se conectar ao banco de dados: %v", err)
	}

	// Verifica se a conexão está ativa.
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("não foi possível verificar a conexão com o banco de dados: %v", err)
	}

	return &Conn{DB: db}, nil
}

// Close fecha a conexão com o banco de dados.
func (c *Conn) Close() error {
	if c.DB != nil {
		return c.DB.Close()
	}
	return errors.New("nenhuma conexão ativa para fechar")
}
