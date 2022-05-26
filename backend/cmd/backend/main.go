package main

import (
	"log"
	"os"

	"awesomeware.org/goblin-wrangler/internal/controllers"
	"awesomeware.org/goblin-wrangler/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func createRouter(app *app) (*gin.Engine, func()) {
	g := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"} // Yeah, this needs to change ;)
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	corsConfig.AddAllowMethods("OPTIONS")
	g.Use(cors.New(corsConfig))

	chat := g.Group("/chat")
	{
		chat.GET("", app.ChatController.FindAll)
		chat.POST("", app.ChatController.Save)
	}

	g.POST("/login", app.AuthController.Login)
	g.GET("/me", app.AuthController.Me)

	return g, func() {}
}

type app struct {
	ChatController controllers.ChatController
	AuthController controllers.AuthController
}

func main() {
	dbPool, err := initDb()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	var chatService services.ChatService = services.NewChatService(dbPool)
	var chatCtrl controllers.ChatController = controllers.NewChatController(chatService)
	var authService services.AuthService = services.NewAuthService(dbPool)
	var authCtrl controllers.AuthController = controllers.NewAuthController(authService)

	var app = &app{
		ChatController: chatCtrl,
		AuthController: authCtrl,
	}

	router, closeable := createRouter(app)
	defer closeable()

	router.Run(":8080")
}
