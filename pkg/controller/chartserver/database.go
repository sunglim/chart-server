package chartserver

import (
	"log"

	"github.com/sunglim/chart-server/internal/database"
	"gorm.io/gorm"
)

var defaultDatabase = CreateDatabase()

func CreateDatabase() *Database {
	db, err := database.CreateDatabase()
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
