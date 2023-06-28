package database

import "github.com/jmoiron/sqlx"

// TODO: Add id column
func Sync(db *sqlx.DB) {
	schema := `
		CREATE TABLE IF NOT EXISTS todo (
			title text,
			category text,
			priority int
		);
	`
	db.MustExec(schema)
}