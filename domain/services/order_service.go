package services

import (
	"context"

	"github.com/lovelyoyrmia/ecommerce/domain/models"
	"github.com/lovelyoyrmia/ecommerce/domain/repository"
	"github.com/lovelyoyrmia/ecommerce/internal/db"
)

type OrderServices interface {
	AddCart(ctx context.Context, req models.CartParams) error
	GetCartProducts(ctx context.Context, req models.CartsParams) (models.CartProducts, error)
}

type orderService struct {
	repo repository.OrderRepositories
}

func NewOrderService(repo repository.OrderRepositories) OrderServices {
	return &orderService{
		repo: repo,
	}
}

// AddCart implements OrderServices.
func (service *orderService) AddCart(ctx context.Context, req models.CartParams) error {
	createCartParams := db.CreateCartTx{
		Uid:      req.Uid,
		Pid:      req.Pid,
		Amount:   req.Amount,
		Quantity: req.Quantity,
	}
	return service.repo.AddCart(ctx, createCartParams)
}

// GetCartProducts implements OrderServices.
func (service *orderService) GetCartProducts(ctx context.Context, req models.CartsParams) (models.CartProducts, error) {
	return service.repo.GetCartProducts(ctx, req)
}
