package repository

import (
	"context"

	"github.com/lovelyoyrmia/ecommerce/domain/models"
	"github.com/lovelyoyrmia/ecommerce/internal/db"
	"github.com/lovelyoyrmia/ecommerce/pkg/utils"
)

type UserRepositories interface {
	CreateUser(ctx context.Context, userParams models.User) (db.User, error)
	LoginUser(ctx context.Context, loginParams models.LoginUser) error
}

type userRepo struct {
	store db.Store
}

func NewUserRepository(store db.Store) UserRepositories {
	return &userRepo{
		store: store,
	}
}

// CreateUser implements UserRepositories.
func (repo *userRepo) CreateUser(ctx context.Context, userParams models.User) (db.User, error) {
	password, err := utils.HashPassword(userParams.Password)
	if err != nil {
		return db.User{}, err
	}
	userParams.Password = password
	up := userParams.ToCreateUserParams()

	return repo.store.CreateUserTx(ctx, up)
}

// LoginUser implements UserRepositories.
func (repo *userRepo) LoginUser(ctx context.Context, loginParams models.LoginUser) error {
	user, err := repo.store.GetUser(ctx, loginParams.Email)
	if err != nil {
		return err
	}

	if err := utils.ComparePassword(user.Password, loginParams.Password); err != nil {
		return err
	}
	return nil
}
