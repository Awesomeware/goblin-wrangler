package controllers

import (
	"net/http"

	"awesomeware.org/goblin-wrangler/internal/models"
	"awesomeware.org/goblin-wrangler/internal/services"
	"github.com/gin-gonic/gin"
)

type ChatController interface {
	FindAll(ctx *gin.Context)
	New(ctx *gin.Context)
}

type chatController struct {
	service services.ChatService
}

func NewChatController(svc services.ChatService) ChatController {
	return &chatController{
		service: svc,
	}
}

func (c *chatController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.FindAll())
}

func (c *chatController) New(ctx *gin.Context) {
	var chatMsg models.ChatMessage
	ctx.BindJSON(&chatMsg)
	c.service.New(chatMsg)
	ctx.JSON(http.StatusOK, chatMsg)
}
