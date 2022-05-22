package main

import (
	"awesomeware.org/kingpin/internal/controllers"
	"awesomeware.org/kingpin/internal/services"
	"github.com/gin-gonic/gin"
)

var (
	chatMsgService services.ChatService       = services.New()
	chatMsgCtrl    controllers.ChatController = controllers.New(chatMsgService)
)

func createRouter() (*gin.Engine, func()) {
	g := gin.Default()

	chat := g.Group("/chat")
	{
		chat.GET("", chatMsgCtrl.FindAll)
		chat.POST("", chatMsgCtrl.Save)
	}

	return g, func() {}
}

func main() {
	router, closeable := createRouter()
	defer closeable()

	router.Run(":8080")
}
