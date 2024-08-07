package orders

import (
	"context"

	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/models"
)

// Usecase - port (первичный)
type Usecase interface {
	// CreateOrder - ...
	CreateOrder(ctx context.Context, orderInfo *CreateOrderInfoDTO) (*models.Order, error)
}

type (
	// WarehouseManagementSystem - port (вторичный)
	WarehouseManagementSystem interface {
		ReserveSKU(ctx context.Context, items []models.SKU) error
	}

	// OrdersRepository - port (вторичный)
	OrdersRepository interface {
		CreateOrder(ctx context.Context, order *models.Order) error
		FindOrders(ctx context.Context, order Filter) ([]*models.Order, error)
	}
)

// Deps -
type Deps struct {
	// Adapters
	WMS              WarehouseManagementSystem
	OrdersRepository OrdersRepository
}

type usecase struct {
	Deps
}

func NewUsecase(d Deps) Usecase {
	return &usecase{Deps: d}
}
