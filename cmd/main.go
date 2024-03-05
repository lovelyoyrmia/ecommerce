package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lovelyoyrmia/ecommerce/domain/handlers"
	"github.com/lovelyoyrmia/ecommerce/domain/middlewares"
	"github.com/lovelyoyrmia/ecommerce/domain/pb"
	"github.com/lovelyoyrmia/ecommerce/domain/repository"
	"github.com/lovelyoyrmia/ecommerce/domain/services"
	"github.com/lovelyoyrmia/ecommerce/internal/db"
	"github.com/lovelyoyrmia/ecommerce/pkg/config"
	"github.com/lovelyoyrmia/ecommerce/pkg/token"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var interruptSignal = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	// Load configuration
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("error occurred : %v", err)
	}

	// Add interrupt signal
	ctx, stop := signal.NotifyContext(context.Background(), interruptSignal...)
	defer stop()

	// Initialize Database
	database := db.NewDatabase(ctx, c)
	store := db.NewStore(database.DB)

	// Initialize token
	maker, err := token.NewPasetoMaker(c)
	if err != nil {
		log.Fatal().Msgf("error occurred : %v", err)
	}

	// Register wait group
	waitGroup, ctx := errgroup.WithContext(ctx)

	// Run Server
	runHTTPGatewayServer(waitGroup, ctx, c, store, maker)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

func runHTTPGatewayServer(
	waitGroup *errgroup.Group,
	ctx context.Context,
	config config.Config,
	store db.Store,
	maker token.Maker,
) {

	// Declare json option for grpc request and response
	jsonOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	// Create grpc serve mux
	grpcMux := runtime.NewServeMux(jsonOpt)

	// Register user services
	userRepo := repository.NewUserRepository(store)
	userService := services.NewUserService(userRepo, maker)
	userGrpc := handlers.NewUserGRPCHandlers(userService)

	err := pb.RegisterUserServiceHandlerServer(ctx, grpcMux, userGrpc)
	if err != nil {
		log.Fatal().Msgf("error occured : %v", err)
	}

	// Register products services
	productRepo := repository.NewProductRepository(store)
	productService := services.NewProductService(productRepo)
	productGrpc := handlers.NewProductGRPCHandlers(productService)

	err = pb.RegisterProductServiceHandlerServer(ctx, grpcMux, productGrpc)
	if err != nil {
		log.Fatal().Msgf("error occured : %v", err)
	}

	// Declare middlewares for grpc
	middle := middlewares.NewMiddlewares(config, maker, store)

	// Register orders services
	orderRepo := repository.NewOrderRepository(store)
	orderService := services.NewOrderService(orderRepo)
	// Inject middlewares to order services
	orderGrpc := handlers.NewOrderGRPCHandlers(orderService, middle)

	err = pb.RegisterOrderServiceHandlerServer(ctx, grpcMux, orderGrpc)
	if err != nil {
		log.Fatal().Msgf("error occured : %v", err)
	}

	// Create http serve mux
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// Register middleware Redoc for API Spec
	ops := middleware.RedocOpts{SpecURL: fmt.Sprintf("%s%s", "./docs/swagger", "/foedie.swagger.json")}
	sh := middleware.Redoc(ops, nil)
	mux.Handle("/docs", sh)

	// Serve API Docs
	fs := http.FileServer(http.Dir("./docs/swagger"))
	mux.Handle(fmt.Sprintf("/docs/swagger%s", "/foedie.swagger.json"), http.StripPrefix("/docs/swagger/", fs))

	// Register http server
	s := &http.Server{
		Addr:         config.GatewayAddress,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	log.Info().Msgf("Starting GRPC server on port: %s", s.Addr)
	// Create go routines to serve http
	waitGroup.Go(func() error {
		err := s.ListenAndServe()
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Fatal().Msgf("cannot start GRPC server: %v", err)
			return err
		}
		return nil
	})

	// Waiting server to gracefully shutdown
	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown HTTP gateway server")

		err := s.Shutdown(context.Background())
		if err != nil {
			log.Error().Msg("failed to shutdown http gateway server")
			return err
		}
		log.Info().Msg("HTTP Gateway server is stopped")
		return nil
	})
}
