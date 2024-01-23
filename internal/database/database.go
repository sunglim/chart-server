package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// CreateDatabase creates the database.
func CreateDatabase() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&HistoricalData{})
	if err != nil {
		return err
	}

	return nil
}
