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
	DeleteCartProduct(ctx context.Context, req models.CartProductParams) error
	CheckoutOrder(ctx context.Context, req models.CartProductParams) (models.Orders, error)
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

// CheckoutOrder implements OrderServices.
func (service *orderService) CheckoutOrder(ctx context.Context, req models.CartProductParams) (models.Orders, error) {
	order, err := service.repo.CheckoutOrder(ctx, req)
	if err != nil {
		return models.Orders{}, err
	}
	orderRes := models.Orders{
		Oid:       order.Oid,
		OrderedAt: order.OrderedAt.Time.String(),
	}
	newOrder := orderRes.ExtractOrderStatus(order.OrderStatus)
	return newOrder, nil
}

// GetCartProducts implements OrderServices.
func (service *orderService) GetCartProducts(ctx context.Context, req models.CartsParams) (models.CartProducts, error) {
	return service.repo.GetCartProducts(ctx, req)
}

// DeleteCartProduct implements OrderServices.
func (service *orderService) DeleteCartProduct(ctx context.Context, req models.CartProductParams) error {
	return service.repo.DeleteCartProduct(ctx, req)
}
