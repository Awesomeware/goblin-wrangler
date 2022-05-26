package services

import (
	"database/sql"
)

type AuthService interface {
	Login() string
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

func (model *authService) Login() string {
	return "OK"
}

func (model *authService) Me() string {
	return "OK"
}
