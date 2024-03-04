package services

import (
	"context"
	"time"

	"github.com/lovelyoyrmia/ecommerce/domain/models"
	"github.com/lovelyoyrmia/ecommerce/domain/repository"
	"github.com/lovelyoyrmia/ecommerce/pkg/token"
)

type UserServices interface {
	CreateUser(ctx context.Context, userParams models.User) (models.UserResponse, error)
	LoginUser(ctx context.Context, loginParams models.LoginUser) (models.LoginResponse, error)
}

type userService struct {
	repo  repository.UserRepositories
	maker token.Maker
}

func NewUserService(repo repository.UserRepositories, maker token.Maker) UserServices {
	return &userService{
		repo:  repo,
		maker: maker,
	}
}

// CreateUser implements UserServices.
func (service *userService) CreateUser(ctx context.Context, userParams models.User) (models.UserResponse, error) {
	user, err := service.repo.CreateUser(ctx, userParams)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		Uid:       user.Uid,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName.String,
	}, nil
}

// LoginUser implements UserServices.
func (service *userService) LoginUser(ctx context.Context, loginParams models.LoginUser) (models.LoginResponse, error) {

	if err := service.repo.LoginUser(ctx, loginParams); err != nil {
		return models.LoginResponse{}, err
	}

	accessToken, payload, errAccessToken := service.maker.GenerateToken(loginParams.Email, time.Hour*24)

	if errAccessToken != nil {
		return models.LoginResponse{}, errAccessToken
	}

	return models.LoginResponse{
		Token: accessToken,
		Email: payload.UID,
	}, nil
}
