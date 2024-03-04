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

type UserGRPCHandlers struct {
	pb.UnimplementedUserServiceServer
	service services.UserServices
}

func NewUserGRPCHandlers(service services.UserServices) *UserGRPCHandlers {
	return &UserGRPCHandlers{
		service: service,
	}
}

func (handler *UserGRPCHandlers) CreateUser(ctx context.Context, createUserParams *pb.CreateUserParams) (*pb.CreateUserResponse, error) {
	userParams := models.User{
		Email:     createUserParams.GetEmail(),
		FirstName: createUserParams.GetFirstName(),
		LastName:  createUserParams.GetLastName(),
		Password:  createUserParams.GetPassword(),
	}
	user, err := handler.service.CreateUser(ctx, userParams)
	if db.ErrorCode(err) == db.UniqueViolation {
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CreateUserResponse{
		Uid:       user.Uid,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (handler *UserGRPCHandlers) LoginUser(ctx context.Context, loginUserParams *pb.LoginUserParams) (*pb.LoginUserResponse, error) {
	loginParams := models.LoginUser{
		Email:    loginUserParams.GetEmail(),
		Password: loginUserParams.GetPassword(),
	}
	loginRes, err := handler.service.LoginUser(ctx, loginParams)
	if errors.Is(err, db.ErrRecordNotFound) {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &pb.LoginUserResponse{
		Token: loginRes.Token,
		Email: loginRes.Email,
	}, nil
}
