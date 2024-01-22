package internal

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"

	"github.com/sunglim/chart-server/internal/database"
	"github.com/sunglim/chart-server/internal/options"
	"github.com/sunglim/chart-server/pkg/strconv"
)

func Wrapper(opts *options.Options) {
	dbfile := "chart.db"
	database.Setup(&dbfile)

	loadHistoricalDatas(opts.HistoryFiles)
}

// Load YahooFinance format csv file.
func loadHistoricalDatas(historicalFiles []string) error {
	for _, file := range historicalFiles {
		stock, err := readFromFile(file)
		if err != nil {
			return err
		}
		// write to db
		writeToDatabase(stock)
	}

	return nil
}

func writeToDatabase(stock *database.Stock) error {
	db, err := sql.Open("sqlite3", "chart.db")
	if err != nil {
		log.Fatalf("Unable to open Sqlite database: %s", err)
	}
	for i := 0; i < len(stock.HistoricalData); i++ {
		database.InsertHistoricalData(db, stock.Name, stock.Symbol, stock.HistoricalData[i].Date, stock.HistoricalData[i].Open, stock.HistoricalData[i].High, stock.HistoricalData[i].Low, stock.HistoricalData[i].Close, stock.HistoricalData[i].AdjClose, stock.HistoricalData[i].Volume)
	}

	defer db.Close()

	return nil
}

func readFromFile(file string) (*database.Stock, error) {
	// read 'file'
	// read csv file
	osFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(osFile)
	records, _ := reader.ReadAll()

	stock := &database.Stock{}
	stock.Name = "MSFT"
	stock.Symbol = "MSFT"

	// skip zero index. it's
	for i := 1; i < len(records); i++ {
		stock.HistoricalData = append(stock.HistoricalData, *recordToYahooFinance(records[i]))
	}

	defer osFile.Close()

	return stock, nil
}

func recordToYahooFinance(record []string) *database.YahooFinance {
	return &database.YahooFinance{Date: record[0],
		Open:     strconv.ParseFloatOrZero(record[1]),
		High:     strconv.ParseFloatOrZero(record[2]),
		Low:      strconv.ParseFloatOrZero(record[3]),
		Close:    strconv.ParseFloatOrZero(record[4]),
		AdjClose: strconv.ParseFloatOrZero(record[5]),
		Volume:   strconv.ParseInt64OrZero(record[6])}
}
