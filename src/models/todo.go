package models

type Todo struct {
	Title    string   `json:"title"`
	Category Category `json:"category"`
	Priority Priority `json:"priority"`
}
