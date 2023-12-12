package model

import "time"

type TodoRequest struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

type Todo struct {
	TodoID      int64     `json:"todo_id"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatorName string    `json:"creator_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type TodoQuery struct {
	Search  *string `form:"search"`
	Page    *uint   `form:"page"`
	PerPage *uint   `form:"per_page"`
}
