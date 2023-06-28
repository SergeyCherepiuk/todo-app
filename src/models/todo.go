package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TODO: All id field
type Todo struct {
	Title    string `json:"title" db:"title"`
	Category string `json:"category" db:"category"`
	Priority int    `json:"priority" db:"priority"`
}

func (t Todo) String() string {
	return fmt.Sprintf("%s, %s, %d\n", t.Title, t.Category, t.Priority)
}

func TodoFromString(s string) (Todo, error) {
	data := strings.Split(s, ", ") // []string{"title", "category", "priority"}
	if len(data) != 3 {
		return Todo{}, errors.New("invalid string")
	}

	priority, err := strconv.ParseInt(data[2], 10, 64)
	if err != nil {
		return Todo{}, err
	}

	return Todo{
		Title:    data[0],
		Category: data[1],
		Priority: int(priority),
	}, nil
}
