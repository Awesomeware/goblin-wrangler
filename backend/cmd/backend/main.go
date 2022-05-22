package main

import (
	"log"
	"os"

	"awesomeware.org/kingpin/internal/controllers"
	"awesomeware.org/kingpin/internal/services"
	"github.com/gin-gonic/gin"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func createRouter(app *app) (*gin.Engine, func()) {
	g := gin.Default()

	chat := g.Group("/chat")
	{
		chat.GET("", app.ChatController.FindAll)
		chat.POST("", app.ChatController.Save)
	}

	return g, func() {}
}

type app struct {
	ChatController controllers.ChatController
}

func main() {
	dbPool, err := initDb()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var chatService services.ChatService = services.New(dbPool)

	var chatCtrl controllers.ChatController = controllers.New(chatService)

	var app = &app{
		ChatController: chatCtrl,
	}

	router, closeable := createRouter(app)
	defer closeable()

	router.Run(":8080")
}
