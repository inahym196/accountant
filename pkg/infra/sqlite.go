package infra

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	driverName = "sqlite3"
)

type SQLiteConnector interface {
	GetConn() *sql.DB
}

type sqliteConnector struct {
	Conn *sql.DB
}

func NewSQLiteConnector(filepath string) SQLiteConnector {
	conn, err := sql.Open(driverName, filepath)
	if err != nil {
		log.Fatal(err)
	}
	return sqliteConnector{conn}
}

func (s sqliteConnector) GetConn() *sql.DB {
	return s.Conn
}
