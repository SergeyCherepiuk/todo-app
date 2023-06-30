package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Sync(db *sqlx.DB) {
	schema := `
		CREATE TABLE IF NOT EXISTS todo (
			id SERIAL PRIMARY KEY,
			title TEXT,
			category TEXT,
			priority INT,
			iscompleted BOOLEAN
		);
	`
	db.MustExec(schema)
	fmt.Println("Database synced successfully")
}
