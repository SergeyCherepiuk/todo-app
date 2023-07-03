package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func MustConnect() *sqlx.DB {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	s := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)
	db := sqlx.MustConnect("postgres", s)
	fmt.Println("Connected to the database successully")
	return db
}
