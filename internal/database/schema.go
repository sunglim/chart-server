package database

import "time"

type HistoricalData struct {
	Symbol    string `gorm:"primaryKey"`
	Date      string `gorm:"primaryKey"`
	Open      float64
	High      float64
	Low       float64
	Close     float64
	AdjClose  float64
	Volume    int64
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}
