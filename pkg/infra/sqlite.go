package infra

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	driverName = "sqlite3"
)

type SQLiteConnector struct {
	Conn *sql.DB
}

func NewSQLiteConnector(filepath string) *SQLiteConnector {
	conn, err := sql.Open(driverName, filepath)
	if err != nil {
		log.Fatal(err)
	}
	return &SQLiteConnector{conn}
}
