package models

type Category struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"category_name"`
}