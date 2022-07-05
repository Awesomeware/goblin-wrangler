package services

import (
	"context"

	"awesomeware.org/goblin-wrangler/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"google.golang.org/api/idtoken"
)

type AuthService interface {
	GetEmailFromPayload(payload idtoken.Payload) string
	ValidateGoogleJWT(jwt models.GoogleSSOToken) (*idtoken.Payload, error)
}

type authService struct {
	DB *pgxpool.Pool
}

func NewAuthService(db *pgxpool.Pool) AuthService {
	return &authService{
		DB: db,
	}
}

func (model *authService) ValidateGoogleJWT(jwt models.GoogleSSOToken) (payload *idtoken.Payload, err error) {
	payload, err = idtoken.Validate(context.Background(), jwt.Credential, viper.GetString("GOOGLE_CLIENT_ID"))
	return
}

func (model *authService) GetEmailFromPayload(payload idtoken.Payload) string {
	return payload.Claims["email"].(string)
}
