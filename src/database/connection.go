package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO: Extract to .env file
const (
	host     = "172.17.0.2"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "todo_app"
)

func Connect() *sqlx.DB {
	s := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sqlx.Open("postgres", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database successully")
	return db
}
