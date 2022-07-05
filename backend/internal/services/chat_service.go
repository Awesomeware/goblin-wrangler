package services

import (
	"awesomeware.org/goblin-wrangler/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ChatService interface {
	New(models.ChatMessage) models.ChatMessage
	FindAll() []models.ChatMessage
}

type chatService struct {
	DB       *pgxpool.Pool
	messages []models.ChatMessage
}

func NewChatService(db *pgxpool.Pool) ChatService {
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
