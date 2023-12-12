package service

import (
	"database/sql"

	"github.com/HtetLinMaung/todo/internal/model"
)

type MessageService struct {
	db *sql.DB
}

func NewMessageService(db *sql.DB) *MessageService {
	return &MessageService{db: db}
}

func (s *MessageService) GetMessage() model.Message {
	return model.Message{Text: "Hello, World!"}
}
