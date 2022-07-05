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
	ValidateGoogleSSOToken(ctx *gin.Context)
	SignupWithGoogle(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthController(authsvc services.AuthService, usersvc services.UserService) AuthController {
	return &authController{
		authService: authsvc,
		userService: usersvc,
	}
}

func (c *authController) SignupWithGoogle(ctx *gin.Context) {
	var req models.GoogleSignupReq
	if err := ctx.Bind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	payload, err := c.authService.ValidateGoogleJWT(req.Token)
	email := c.authService.GetEmailFromPayload(*payload)
	username := req.Username

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = c.userService.GetUserByEmail(email)
	if err != nil {
		// TODO: Probably don't error on this -- maybe just skip straight to creating JWT token?
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, err := c.userService.CreateUser(email, username)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// TODO: Create and return a JWT token, rather than just a user
	ctx.JSON(http.StatusOK, id)
}

func (c *authController) ValidateGoogleSSOToken(ctx *gin.Context) {
	var token models.GoogleSSOTokenReq
	if err := ctx.Bind(&token); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := c.authService.ValidateGoogleJWT(token)
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
