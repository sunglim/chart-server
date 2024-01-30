package csvinsert

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sunglim/chart-server/internal/database"
	"github.com/sunglim/chart-server/pkg/strconv"
)

func fromCSVRecord(symbol string, colums []string) database.HistoricalData {
	return database.HistoricalData{
		Symbol:   symbol,
		Date:     colums[0],
		Open:     strconv.ParseFloat64OrZero(colums[1]),
		High:     strconv.ParseFloat64OrZero(colums[2]),
		Low:      strconv.ParseFloat64OrZero(colums[3]),
		Close:    strconv.ParseFloat64OrZero(colums[4]),
		AdjClose: strconv.ParseFloat64OrZero(colums[5]),
		Volume:   strconv.ParseInt64OrZero(colums[6]),
	}
}

func fromCSV(reader io.Reader) ([]database.HistoricalData, error) {
	var data []database.HistoricalData
	i := 0
	r := csv.NewReader(reader)
	for {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return data, fmt.Errorf("%w", err)
		}

		// Skip first line.
		if i == 0 {
			i++
			continue
		}

		// TODO: Pass symbol from file name. or argument?
		HistoricalData := fromCSVRecord("MSFT", record)
		data = append(data, HistoricalData)
	}

	return data, nil
}

func FromCSVFile(filePath string) ([]database.HistoricalData, error) {
	// TODO: if the given path is already abs path?

	absPath, _ := filepath.Abs(filePath)

	osFile, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer osFile.Close()

	return fromCSV(osFile)
}
