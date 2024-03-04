package repository

import (
	"context"

	"github.com/lovelyoyrmia/ecommerce/domain/models"
	"github.com/lovelyoyrmia/ecommerce/internal/db"
)

type ProductRepositories interface {
	GetProducts(ctx context.Context, productsParams models.ProductsParams) ([]db.GetProductsRow, error)
	GetProductDetails(ctx context.Context, productParams models.ProductDetailsParams) (db.GetProductDetailsRow, error)
	GetProductsByCategory(ctx context.Context, category string) ([]db.GetProductsByCategoryRow, error)
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
	var limit int32
	var offset int32
	if productsParams.Limit != nil {
		limit = *productsParams.Limit
	} else {
		limit = 5
	}

	if productsParams.Offset != nil {
		offset = *productsParams.Offset
	} else {
		offset = 0
	}

	return repo.store.GetProducts(ctx, db.GetProductsParams{
		Limit:  limit,
		Offset: offset,
	})
}

// GetProductsByCategory implements ProductRepositories.
func (repo *productRepo) GetProductsByCategory(ctx context.Context, category string) ([]db.GetProductsByCategoryRow, error) {
	categoryName, err := repo.store.GetProductCategory(ctx, category)
	if err != nil {
		return []db.GetProductsByCategoryRow{}, err
	}

	return repo.store.GetProductsByCategory(ctx, categoryName)
}

// GetProductDetails implements ProductRepositories.
func (repo *productRepo) GetProductDetails(ctx context.Context, productParams models.ProductDetailsParams) (db.GetProductDetailsRow, error) {
	return repo.store.GetProductDetails(ctx, productParams.Pid)
}
