package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// CreateDatabase creates the database.
func CreateDatabase() (*gorm.DB, error) {
	return CreateDatabaseFromFilePath("./database.db")
}

func CreateDatabaseFromFilePath(filePath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&HistoricalData{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// InsertHistoricalData inserts historical data into the database.
func InsertHistoricalData(db *gorm.DB, data *HistoricalData) {
	db.Create(data)
}

func InsertHistoricalDatas(db *gorm.DB, datas []HistoricalData) {
	db.CreateInBatches(datas, 100)
}

func GetHistoricalDatas(db *gorm.DB, symbol string) []HistoricalData {
	var datas []HistoricalData
	db.Find(&datas, "symbol = ?", symbol)
	return datas
}

func GetClosingPrices(db *gorm.DB, symbol string) []float64 {
	var datas = GetHistoricalDatas(db, symbol)
	var prices []float64
	for _, data := range datas {
		prices = append(prices, data.Close)
	}
	return prices
}
