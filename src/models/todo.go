package models

import (
	"fmt"
)

type Todo struct {
	ID          uint64 `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Category    string `json:"category" db:"category"`
	Priority    int    `json:"priority" db:"priority"`
	IsCompleted bool   `json:"isCompleted" db:"iscompleted"`
}

func (t Todo) String() string {
	return fmt.Sprintf("%s, %s, %d\n", t.Title, t.Category, t.Priority)
}
