package main

import (
	"context"
	"log"

	ordersrepository "github.com/moguchev/microservices_courcse/orders_management_system/internal/app/adapters/orders_repository"
	warehousemanagementsystem "github.com/moguchev/microservices_courcse/orders_management_system/internal/app/adapters/warehouse_management_system"
	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/server"
	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/usecase/orders"
	middleware "github.com/moguchev/microservices_courcse/orders_management_system/internal/middleware/errors"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// =========================
	// adapters
	// =========================

	// repository

	ordersRepo := ordersrepository.NewRepository()

	// services

	wmsAdapter := warehousemanagementsystem.NewClient()

	// =========================
	// usecases
	// =========================

	ordersUsecase := orders.NewUsecase(orders.Deps{
		WMS:              wmsAdapter,
		OrdersRepository: ordersRepo,
	})

	// =========================
	// delivery
	// =========================

	config := server.Config{
		GRPCPort:        ":8082",
		GRPCGatewayPort: ":8080",
		ChainUnaryInterceptors: []grpc.UnaryServerInterceptor{
			middleware.ErrorsUnaryInterceptor(),
		},
	}

	srv, err := server.New(ctx, config, server.Deps{
		OrdersUsecase: ordersUsecase,
		// Dependency injection
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("run: %v", err)
	}
}
