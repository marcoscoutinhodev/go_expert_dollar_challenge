package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetSqliteConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "goexpert.db")

	if err != nil {
		panic(err)
	}

	return db
}
