package handlers

import (
	"context"
	"errors"

	"github.com/lovelyoyrmia/ecommerce/domain/middlewares"
	"github.com/lovelyoyrmia/ecommerce/domain/models"
	"github.com/lovelyoyrmia/ecommerce/domain/pb"
	"github.com/lovelyoyrmia/ecommerce/domain/services"
	"github.com/lovelyoyrmia/ecommerce/internal/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderGRPCHandlers struct {
	pb.UnimplementedOrderServiceServer
	service    services.OrderServices
	middleware *middlewares.Middlewares
}

func NewOrderGRPCHandlers(service services.OrderServices, middleware *middlewares.Middlewares) *OrderGRPCHandlers {
	return &OrderGRPCHandlers{
		service:    service,
		middleware: middleware,
	}
}

func (handler *OrderGRPCHandlers) AddCart(ctx context.Context, req *pb.CreateCartParams) (*pb.CreateCartResponse, error) {
	user, err := handler.middleware.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "not authenticated")
	}

	err = handler.service.AddCart(ctx, models.CartParams{
		Uid:      user.Uid,
		Pid:      req.GetPid(),
		Quantity: req.GetQuantity(),
		Amount:   int32(req.GetAmount()),
	})

	if db.ErrorCode(err) == db.UniqueViolation {
		return nil, status.Error(codes.AlreadyExists, "product already exists")
	}

	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &pb.CreateCartResponse{
		Message: "Product added to cart",
	}, nil
}

func (handler *OrderGRPCHandlers) GetCarts(ctx context.Context, req *pb.GetCartUserParams) (*pb.GetCartUserResponse, error) {
	user, err := handler.middleware.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "not authenticated")
	}
	carts, err := handler.service.GetCartProducts(ctx, models.CartsParams{
		Uid: user.Uid,
	})

	if errors.Is(err, db.ErrRecordNotFound) {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	var productsRes []*pb.CartProduct

	for _, v := range carts.Products {
		productsRes = append(productsRes, &pb.CartProduct{
			Pid:         v.Pid,
			Name:        v.Name,
			Description: v.Description,
			Category:    v.CategoryName,
			Stock:       v.Stock,
			Price:       int64(v.Price),
			Quantity:    v.Quantity,
			Amount:      int64(v.Amount),
		})
	}

	return &pb.GetCartUserResponse{
		Oid:      carts.Oid,
		Products: productsRes,
	}, nil
}
