package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

const (
	dbDriver     = "mysql"
	dbUser       = "root"
	dbPassword   = "1234"
	dbHost       = "localhost"
	dbPort       = "3306"
	dbName       = "loja"
	maxWaitTime  = 60 * time.Second
	pingInterval = 5 * time.Second
)

var Store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

type Config interface {
	Getdb() *sql.DB
}

type Conn struct {
	DB    *sql.DB
	store sessions.CookieStore
}

func (db Conn) Getdb() *sql.DB {
	return db.DB
}

func (db *Conn) GetStore() *sessions.CookieStore {
	return &db.store
}

func NewConn() (Config, error) {
	// Cria uma conexão com o banco de dados.
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dbURI)
	db, err := sql.Open(dbDriver, dbURI)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	// Esperar até que o MySQL esteja pronto
	waitForDB(db)

	return &Conn{DB: db}, nil
}

func waitForDB(db *sql.DB) {
	deadline := time.Now().Add(maxWaitTime)

	for {
		err := db.Ping()
		if err == nil {
			log.Println("Banco de dados pronto para conexões")
			return
		}

		if time.Now().After(deadline) {
			log.Fatalf("Timeout esperando pelo banco de dados: %v", err)
		}

		log.Printf("Aguardando pelo banco de dados... %v", err)
		time.Sleep(pingInterval)
	}
}
