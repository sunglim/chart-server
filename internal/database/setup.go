// Initial settings to use db. Such as creating tables.
package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Setup(dbFile *string) {
	// Create the db "connection"
	db, err := sql.Open("sqlite3", *dbFile)
	if err != nil {
		log.Fatalf("Unable to open Sqlite database: %s", err)
	}
	defer db.Close()

	CreateDatabase(db)
}
