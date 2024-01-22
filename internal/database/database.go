// Sql layer. No logical change added.

package database

import "database/sql"

func CreateDatabase(db *sql.DB) error {
	sqlStmt := `
	create table if not exists historicaldata (symbol text, name text, date text, open real, high real, low real, close real, adj_close real, volume integer);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil
}

func InsertHistoricalData(db *sql.DB, symbol string, name string, date string, open float64, high float64, low float64, close float64, adj_close float64, volume int64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into historicaldata(symbol, name, date, open, high, low, close, adj_close, volume) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(symbol, name, date, open, high, low, close, adj_close, volume)
	if err != nil {
		return err
	}

	defer stmt.Close()

	tx.Commit()

	return nil
}
