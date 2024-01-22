package database

// Yahoo Finance struct
type YahooFinance struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	AdjClose float64
	Volume   int64
}

type Stock struct {
	Symbol         string
	Name           string
	HistoricalData []YahooFinance
}
