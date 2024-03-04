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
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("error occurred : %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignal...)
	defer stop()

	database := db.NewDatabase(ctx, c)
	store := db.NewStore(database.DB)

	maker, err := token.NewPasetoMaker(c)
	if err != nil {
		log.Fatal().Msgf("error occurred : %v", err)
	}

	waitGroup, ctx := errgroup.WithContext(ctx)

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

	jsonOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOpt)

	userRepo := repository.NewUserRepository(store)
	userService := services.NewUserService(userRepo, maker)
	userGrpc := handlers.NewUserGRPCHandlers(userService)

	err := pb.RegisterUserServiceHandlerServer(ctx, grpcMux, userGrpc)
	if err != nil {
		log.Fatal().Msgf("error occured : %v", err)
	}

	productRepo := repository.NewProductRepository(store)
	productService := services.NewProductService(productRepo)
	productGrpc := handlers.NewProductGRPCHandlers(productService)

	err = pb.RegisterProductServiceHandlerServer(ctx, grpcMux, productGrpc)
	if err != nil {
		log.Fatal().Msgf("error occured : %v", err)
	}

	middle := middlewares.NewMiddlewares(config, maker, store)

	orderRepo := repository.NewOrderRepository(store)
	orderService := services.NewOrderService(orderRepo)
	orderGrpc := handlers.NewOrderGRPCHandlers(orderService, middle)

	err = pb.RegisterOrderServiceHandlerServer(ctx, grpcMux, orderGrpc)
	if err != nil {
		log.Fatal().Msgf("error occured : %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	ops := middleware.RedocOpts{SpecURL: fmt.Sprintf("%s%s", "./docs/swagger", "/foedie.swagger.json")}
	sh := middleware.Redoc(ops, nil)
	mux.Handle("/docs", sh)

	fs := http.FileServer(http.Dir("./docs/swagger"))
	mux.Handle(fmt.Sprintf("/docs/swagger%s", "/foedie.swagger.json"), http.StripPrefix("/docs/swagger/", fs))

	s := &http.Server{
		Addr:         config.GatewayAddress,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	log.Info().Msgf("Starting GRPC server on port: %s", s.Addr)
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
