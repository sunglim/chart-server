package chartserver

import (
	"log"

	"github.com/sunglim/chart-server/internal/database"
	"gorm.io/gorm"
)

var defaultDatabase = CreateDatabase("./database.db")

func CreateDatabase(path string) *Database {
	db, err := database.CreateDatabaseFromFilePath(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Database initialized")

	return &Database{db: db}
}

type Database struct {
	db *gorm.DB
}

func Default() *Database {
	return defaultDatabase
}
