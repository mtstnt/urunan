package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Singleton for now.
var (
	Q  Querier
	db *sql.DB
)

func Recreate() {
	query, err := os.ReadFile("./sql/schema.sql")
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec(string(query)); err != nil {
		panic(err)
	}
}

func Initialize() {
	_db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}
	db = _db
	Q = New(_db)
}
