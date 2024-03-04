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
	products, err := handler.service.GetProducts(ctx, models.ProductsParams{
		Limit:  int32(productsParams.Limit),
		Offset: int32(productsParams.Offset),
	})
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	var productsRes []*pb.Product

	for _, v := range products.Products {
		productsRes = append(productsRes, &pb.Product{
			Pid:         v.Pid,
			Name:        v.Name,
			Description: v.Description.String,
			Category:    v.CategoryName.String,
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
