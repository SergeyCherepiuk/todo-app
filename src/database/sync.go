package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// TODO: Add id column
func Sync(db *sqlx.DB) {
	schema := `
		CREATE TABLE IF NOT EXISTS todo (
			id SERIAL PRIMARY KEY,
			title TEXT,
			category TEXT,
			priority INT
		);
	`
	if _, err := db.Exec(schema); err != nil {
		panic(err)
	}
	fmt.Println("Database synced successfully")
}
