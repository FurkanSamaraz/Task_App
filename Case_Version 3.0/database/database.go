package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "172754"
	dbname   = "postgres"
)

func Openconnention() *sql.DB {

	psq := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psq)

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
