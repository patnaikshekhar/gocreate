package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	conn *sql.DB
}

func New(relationDBToUse, connectionString string) (*Db, error) {
	conn, err := sql.Open(relationDBToUse, connectionString)
	if err != nil {
		return nil, err
	}
	return &Db{
		conn: conn,
	}, nil
}
