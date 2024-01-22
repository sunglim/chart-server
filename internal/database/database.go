package database

import "database/sql"

func CreateDatabase(db *sql.DB) error {
	sqlStmt := `
	create table if not exists nodes (symbol text, name text, date text, open real, high real, low real, close real, adj_close real, volume integer);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil
}
