package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// CreateDatabase creates the database.
func CreateDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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

func GetHistoricalDatas(db *gorm.DB) []HistoricalData {
	var datas []HistoricalData
	db.Find(&datas)
	return datas
}

func GetClosingPrices(db *gorm.DB) []float64 {
	var datas = GetHistoricalDatas(db)
	var prices []float64
	for _, data := range datas {
		prices = append(prices, data.Close)
	}
	return prices
}
