package internal

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/sunglim/chart-server/internal/database"
	"github.com/sunglim/chart-server/internal/options"
)

func Wrapper(opts *options.Options) {
	dbfile := "chart.db"
	database.Setup(&dbfile)

	loadHistoricalDatas(opts.HistoryFiles)
}

// Load YahooFinance format csv file.
func loadHistoricalDatas(historicalFiles []string) error {
	for _, file := range historicalFiles {
		_, err := readFromFile(file)
		if err != nil {
			return err
		}
	}

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

	for _, eachrecord := range records {
		fmt.Println(eachrecord)
	}

	defer osFile.Close()

	return nil, nil
}
