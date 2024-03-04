package repository

import (
	"context"

	"github.com/lovelyoyrmia/ecommerce/domain/models"
	"github.com/lovelyoyrmia/ecommerce/internal/db"
)

type ProductRepositories interface {
	GetProducts(ctx context.Context, productsParams models.ProductsParams) ([]db.GetProductsRow, error)
	GetProductDetails(ctx context.Context, productParams models.ProductDetailsParams) (db.GetProductDetailsRow, error)
}

type productRepo struct {
	store db.Store
}

func NewProductRepository(store db.Store) ProductRepositories {
	return &productRepo{
		store: store,
	}
}

// GetProducts implements ProductRepositories.
func (repo *productRepo) GetProducts(ctx context.Context, productsParams models.ProductsParams) ([]db.GetProductsRow, error) {
	products, err := repo.store.GetProducts(ctx, db.GetProductsParams{
		Limit:  productsParams.Limit,
		Offset: productsParams.Offset,
	})
	if err != nil {
		return []db.GetProductsRow{}, err
	}

	return products, nil
}

// GetProductDetails implements ProductRepositories.
func (repo *productRepo) GetProductDetails(ctx context.Context, productParams models.ProductDetailsParams) (db.GetProductDetailsRow, error) {
	return repo.store.GetProductDetails(ctx, productParams.Pid)
}
