package services

import (
	"database/sql"

	"awesomeware.org/goblin-wrangler/internal/models"
)

type ChatService interface {
	New(models.ChatMessage) models.ChatMessage
	FindAll() []models.ChatMessage
}

type chatService struct {
	DB       *sql.DB
	messages []models.ChatMessage
}

func NewChatService(db *sql.DB) ChatService {
	return &chatService{
		DB: db,
	}
}

func (model *chatService) New(msg models.ChatMessage) models.ChatMessage {
	model.messages = append(model.messages, msg)

	return msg
}

func (model *chatService) FindAll() []models.ChatMessage {
	return model.messages
}
