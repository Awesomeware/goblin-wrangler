package main

import (
	"fmt"
	"log"
	"os"

	"awesomeware.org/goblin-wrangler/internal/controllers"
	"awesomeware.org/goblin-wrangler/internal/db"
	"awesomeware.org/goblin-wrangler/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setupViper(env string) {
	viper.SetDefault("PORT", "8080")
	viper.SetEnvPrefix("vpr")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName(fmt.Sprintf("viper.%s", env))
	viper.SetConfigType("dotenv")
	viper.ReadInConfig()
}

func getCorsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{viper.GetString("FRONTEND_CORS_ORIGIN")}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	corsConfig.AddAllowMethods("OPTIONS")
	return corsConfig
}

func createRouter(app *app) (*gin.Engine, func()) {
	g := gin.Default()
	g.Use(cors.New(getCorsConfig()))

	chat := g.Group("/chat")
	{
		chat.GET("", app.ChatController.FindAll)
		chat.POST("", app.ChatController.New)
	}

	g.POST("/auth/google", app.AuthController.ValidateGoogleSSOToken)

	return g, func() {}
}

type app struct {
	ChatController controllers.ChatController
	AuthController controllers.AuthController
}

func main() {
	setupViper("local")

	dbPool, err := db.New()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	var chatService services.ChatService = services.NewChatService(dbPool)
	var userService services.UserService = services.NewUserService(dbPool)
	var authService services.AuthService = services.NewAuthService(dbPool)
	// TODO: JWT service with GenerateToken() / ValidateToken() for our own JWTs

	var chatCtrl controllers.ChatController = controllers.NewChatController(chatService)
	var authCtrl controllers.AuthController = controllers.NewAuthController(authService, userService)

	var app = &app{
		ChatController: chatCtrl,
		AuthController: authCtrl,
	}

	router, closeable := createRouter(app)
	defer closeable()

	router.Run(":" + viper.GetString("PORT"))
}
