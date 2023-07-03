package models

type Todo struct {
	ID          uint64 `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Priority    int    `json:"priority" db:"priority"`
	IsCompleted bool   `json:"is_completed" db:"is_completed"`
	UserID      uint64 `json:"user_id" db:"user_id"`
	CategoryID    uint64 `json:"category_id" db:"category_id"`
}
