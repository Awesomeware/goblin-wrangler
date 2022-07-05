package services

import (
	"context"
	"database/sql"
	"fmt"

	"awesomeware.org/goblin-wrangler/internal/models"
	"github.com/spf13/viper"
	"google.golang.org/api/idtoken"
)

type AuthService interface {
	Login() string
	ValidateGoogleSSOToken(token models.GoogleSSOTokenReq) (models.GoogleSSOTokenClaims, error)
	Me() string
}

type authService struct {
	DB *sql.DB
}

func NewAuthService(db *sql.DB) AuthService {
	return &authService{
		DB: db,
	}
}

func (model *authService) ValidateGoogleSSOToken(token models.GoogleSSOTokenReq) (models.GoogleSSOTokenClaims, error) {
	payload, err := idtoken.Validate(context.Background(), token.Credential, viper.GetString("GOOGLE_CLIENT_ID"))

	if err != nil {
		return models.GoogleSSOTokenClaims{}, err
	}

	fmt.Println("PAYLOAD:")
	fmt.Println(payload)
	fmt.Println("PAYLOAD OVER:")

	return models.GoogleSSOTokenClaims{
		Claims: payload.Claims,
	}, nil
}

func (model *authService) Login() string {
	return "OK"
}

func (model *authService) Me() string {
	return "OK"
}
