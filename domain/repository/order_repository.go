package repository

import (
	"context"

	"github.com/lovelyoyrmia/ecommerce/domain/models"
	"github.com/lovelyoyrmia/ecommerce/internal/db"
)

type OrderRepositories interface {
	AddCart(ctx context.Context, createCartParams db.CreateCartTx) error
	GetCartProducts(ctx context.Context, req models.CartsParams) (models.CartProducts, error)
}

type orderRepo struct {
	store db.Store
}

func NewOrderRepository(store db.Store) OrderRepositories {
	return &orderRepo{
		store: store,
	}
}

// AddCart implements OrderRepositories.
func (repo *orderRepo) AddCart(ctx context.Context, createCartParams db.CreateCartTx) error {
	return repo.store.CreateCartUserTx(ctx, createCartParams)
}

// GetCartProducts implements OrderRepositories.
func (repo *orderRepo) GetCartProducts(ctx context.Context, req models.CartsParams) (models.CartProducts, error) {
	order, err := repo.store.GetCartUser(ctx, req.Uid)
	if err != nil {
		return models.CartProducts{}, err
	}

	orderItems, err := repo.store.GetCartProducts(ctx, order.Oid)
	if err != nil {
		return models.CartProducts{}, err
	}

	var products []models.CartProduct
	for _, v := range orderItems {
		product, err := repo.store.GetProductDetails(ctx, v.Pid)
		if err != nil {
			return models.CartProducts{}, err
		}

		products = append(products, models.CartProduct{
			Pid:          v.Pid,
			Name:         product.Name,
			Description:  product.Description.String,
			CategoryName: product.CategoryName.String,
			Stock:        product.Stock,
			Price:        int64(product.Price),
			Quantity:     v.Quantity,
			Amount:       int64(v.Quantity * v.Amount),
		})
	}

	return models.CartProducts{
		Oid:      order.Oid,
		Products: products,
	}, nil
}
