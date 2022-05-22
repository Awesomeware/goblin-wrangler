package controllers

import (
	"net/http"

	"awesomeware.org/kingpin/internal/models"
	"awesomeware.org/kingpin/internal/services"
	"github.com/gin-gonic/gin"
)

type ChatController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
}

type controller struct {
	service services.ChatService
}

func New(svc services.ChatService) ChatController {
	return &controller{
		service: svc,
	}
}

func (c *controller) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.FindAll())
}

func (c *controller) Save(ctx *gin.Context) {
	var chatMsg models.ChatMessage
	ctx.BindJSON(&chatMsg)
	c.service.Save(chatMsg)
	ctx.JSON(http.StatusOK, chatMsg)
}
