package database

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Singleton for now.
var (
	Q  Querier
	DB *sql.DB
)

func Recreate() {
	query, err := os.ReadFile("./sql/schema.sql")
	if err != nil {
		panic(err)
	}

	if _, err := DB.Exec(string(query)); err != nil {
		panic(err)
	}

	slog.Info("Successfully refreshed database schema and data...")
}

func Initialize() {
	_db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}
	DB = _db
	Q = New(_db)

	slog.Info("Successfully opened database connection")
}
