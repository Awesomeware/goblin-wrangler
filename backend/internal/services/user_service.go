package services

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserService interface {
	CreateUser(email string, username string) (id int, err error)
	GetUserByEmail(email string) (id int, err error)
}

type userService struct {
	DB *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) UserService {
	return &userService{
		DB: db,
	}
}

func (model *userService) GetUserByEmail(email string) (id int, err error) {
	err = model.DB.QueryRow(context.Background(), "SELECT id FROM users WHERE email=$1", email).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Checking for existing user failed: %v\n", err)
		return -1, err
	}

	// TODO: return actual user, rather than just id
	return
}

func (model *userService) CreateUser(email string, username string) (id int, err error) {
	stmt := `
	INSERT INTO users (email, username)
	VALUES ($1, $2)
	RETURNING id
	`
	err = model.DB.QueryRow(context.Background(), stmt, email, username).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Inserting new user failed: %v\n", err)
		return
	}

	fmt.Println("New user ID is: ", id)
	// TODO: return actual user, rather than just id
	return
}
