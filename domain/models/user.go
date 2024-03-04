package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lovelyoyrmia/ecommerce/internal/db"
)

type User struct {
	Uid       string
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func (user User) ToCreateUserParams() db.CreateUserParams {

	return db.CreateUserParams{
		Uid:       uuid.NewString(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName: pgtype.Text{
			String: user.LastName,
			Valid:  true,
		},
		Password: user.Password,
	}
}

type UserResponse struct {
	Uid       string
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
}
