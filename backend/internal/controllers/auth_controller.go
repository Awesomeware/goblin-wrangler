package controllers

import (
	"net/http"

	"awesomeware.org/goblin-wrangler/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Me(ctx *gin.Context)
}

type authController struct {
	service services.AuthService
}

func NewAuthController(svc services.AuthService) AuthController {
	return &authController{
		service: svc,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.Login())
}

func (c *authController) Me(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.Me())
}
