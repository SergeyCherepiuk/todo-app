package models

type Todo struct {
	ID          uint64 `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Category    uint64 `json:"category_id" db:"category_id"`
	Priority    int    `json:"priority" db:"priority"`
	IsCompleted bool   `json:"is_completed" db:"is_completed"`
}
