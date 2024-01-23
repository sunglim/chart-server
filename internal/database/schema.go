package database

type HistoricalData struct {
	Symbol   string
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	AdjClose float64
	Volume   int64
}
