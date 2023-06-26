package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Todo struct {
	Title    string   `json:"title"`
	Category Category `json:"category"`
	Priority Priority `json:"priority"`
}

func (t Todo) String() string {
	return fmt.Sprintf("%s, %s, %d\n", t.Title, t.Category.Name, t.Priority)
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
		Title: data[0],
		Category: Category{Name: data[1]},
		Priority: Priority(priority),
	}, nil
}