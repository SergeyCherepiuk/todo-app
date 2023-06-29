package models

import (
	"fmt"
)

type Todo struct {
	ID       string `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Category string `json:"category" db:"category"`
	Priority int    `json:"priority" db:"priority"`
}

func (t Todo) String() string {
	return fmt.Sprintf("%s, %s, %d\n", t.Title, t.Category, t.Priority)
}
