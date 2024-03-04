package handlers

import (
	"context"
	"errors"

	"github.com/lovelyoyrmia/ecommerce/domain/models"
	"github.com/lovelyoyrmia/ecommerce/domain/pb"
	"github.com/lovelyoyrmia/ecommerce/domain/services"
	"github.com/lovelyoyrmia/ecommerce/internal/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductGRPCHandlers struct {
	pb.UnimplementedProductServiceServer
	service services.ProductServices
}

func NewProductGRPCHandlers(service services.ProductServices) *ProductGRPCHandlers {
	return &ProductGRPCHandlers{
		service: service,
	}
}

func (handler *ProductGRPCHandlers) GetProducts(ctx context.Context, productsParams *pb.GetProductParams) (*pb.GetProductResponse, error) {

	var products []models.Product
	if productsParams.Category != nil {
		newProducts, err := handler.service.GetProductsByCategory(ctx, productsParams.GetCategory())
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "category not found")
		}

		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}

		products = newProducts.Products
	} else {
		newProducts, err := handler.service.GetProducts(ctx, models.ProductsParams{
			Limit:  productsParams.Limit,
			Offset: productsParams.Offset,
		})
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		products = newProducts.Products
	}

	var productsRes []*pb.Product

	for _, v := range products {
		productsRes = append(productsRes, &pb.Product{
			Pid:         v.Pid,
			Name:        v.Name,
			Description: v.Description,
			Category:    v.Category,
			Stock:       int64(v.Stock),
			Price:       int64(v.Price),
		})
	}

	return &pb.GetProductResponse{
		Products: productsRes,
	}, nil
}

func (handler *ProductGRPCHandlers) GetProductDetails(ctx context.Context, productsParams *pb.GetProductDetailsParams) (*pb.Product, error) {
	product, err := handler.service.GetProductDetails(ctx, models.ProductDetailsParams{
		Pid: productsParams.Pid,
	})

	if errors.Is(err, db.ErrRecordNotFound) {
		return nil, status.Error(codes.NotFound, "product not found")
	}

	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &pb.Product{
		Pid:         product.Pid,
		Name:        product.Name,
		Description: product.Description,
		Stock:       int64(product.Stock),
		Price:       int64(product.Price),
		Category:    product.Category,
	}, nil
}
