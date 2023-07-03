package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func MustSync(db *sqlx.DB) {
	schema, err := os.ReadFile("./database/schema.sql")
	if err != nil {
		panic(err)
	}
	db.MustExec(string(schema))
	fmt.Println("Database synced successfully")
}
