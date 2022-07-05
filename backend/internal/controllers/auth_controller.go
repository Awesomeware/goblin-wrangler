package controllers

import (
	"net/http"
	"net/url"

	"awesomeware.org/goblin-wrangler/internal/models"
	"awesomeware.org/goblin-wrangler/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Me(ctx *gin.Context)
	ValidateGoogleSSOToken(ctx *gin.Context)
}

type authController struct {
	service services.AuthService
}

func NewAuthController(svc services.AuthService) AuthController {
	return &authController{
		service: svc,
	}
}

func (c *authController) ValidateGoogleSSOToken(ctx *gin.Context) {
	var token models.GoogleSSOTokenReq
	if err := ctx.Bind(&token); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := c.service.ValidateGoogleSSOToken(token)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	location, err := url.Parse(viper.GetString("FRONTEND_CORS_ORIGIN") + "/api/sso_callback")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	ctx.SetCookie("g_sso_jwt", token.Credential, 5, "/", ctx.Request.URL.Hostname(), false, false)

	q := url.Values{}
	q.Set("email", resp.Claims["email"].(string))
	q.Set("name", resp.Claims["name"].(string))
	q.Set("sub", resp.Claims["sub"].(string))
	location.RawQuery = q.Encode()
	//location := url.URL{Host: viper.GetString("FRONTEND_CORS_ORIGIN"), Path: "/api/sso_callback", RawQuery: q.Encode()}
	ctx.Redirect(http.StatusFound, location.String())
	//ctx.JSON(http.StatusTemporaryRedirect, resp)
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.Login())
}

func (c *authController) Me(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.Me())
}
