package services

import (
	"context"

	"github.com/lovelyoyrmia/ecommerce/domain/models"
	"github.com/lovelyoyrmia/ecommerce/domain/repository"
)

type ProductServices interface {
	GetProducts(ctx context.Context, productsParams models.ProductsParams) (models.ProductsResponse, error)
	GetProductsByCategory(ctx context.Context, category string) (models.ProductsResponse, error)
	GetProductDetails(ctx context.Context, productParams models.ProductDetailsParams) (models.Product, error)
}

type productService struct {
	repo repository.ProductRepositories
}

func NewProductService(repo repository.ProductRepositories) ProductServices {
	return &productService{
		repo: repo,
	}
}

// GetProducts implements ProductServices.
func (service *productService) GetProducts(ctx context.Context, productsParams models.ProductsParams) (models.ProductsResponse, error) {

	products, err := service.repo.GetProducts(ctx, productsParams)
	if err != nil {
		return models.ProductsResponse{}, err
	}

	var newProducts []models.Product
	for _, v := range products {
		newProducts = append(newProducts, models.Product{
			Pid:         v.Pid,
			Name:        v.Name,
			Category:    v.CategoryName.String,
			Description: v.Description.String,
			Price:       v.Price,
			Stock:       v.Stock,
		})
	}

	return models.ProductsResponse{
		Products: newProducts,
	}, nil
}

// GetProductsByCategory implements ProductServices.
func (service *productService) GetProductsByCategory(ctx context.Context, category string) (models.ProductsResponse, error) {
	products, err := service.repo.GetProductsByCategory(ctx, category)
	if err != nil {
		return models.ProductsResponse{}, err
	}

	var newProducts []models.Product
	for _, v := range products {
		newProducts = append(newProducts, models.Product{
			Pid:         v.Pid,
			Name:        v.Name,
			Category:    v.CategoryName.String,
			Description: v.Description.String,
			Price:       v.Price,
			Stock:       v.Stock,
		})
	}

	return models.ProductsResponse{
		Products: newProducts,
	}, nil
}

// GetProductDetails implements ProductServices.
func (service *productService) GetProductDetails(ctx context.Context, productParams models.ProductDetailsParams) (models.Product, error) {
	product, err := service.repo.GetProductDetails(ctx, productParams)
	if err != nil {
		return models.Product{}, err
	}
	return models.Product{
		Pid:         product.Pid,
		Name:        product.Name,
		Description: product.Description.String,
		Category:    product.CategoryName.String,
		Price:       product.Price,
		Stock:       product.Stock,
	}, nil
}
