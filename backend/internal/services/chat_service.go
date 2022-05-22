package services

import "awesomeware.org/kingpin/internal/models"

type ChatService interface {
	Save(models.ChatMessage) models.ChatMessage
	FindAll() []models.ChatMessage
}

type chatService struct {
	messages []models.ChatMessage
}

func New() ChatService {
	return &chatService{}
}

func (model *chatService) Save(msg models.ChatMessage) models.ChatMessage {
	model.messages = append(model.messages, msg)
	return msg
}

func (model *chatService) FindAll() []models.ChatMessage {
	return model.messages
}
