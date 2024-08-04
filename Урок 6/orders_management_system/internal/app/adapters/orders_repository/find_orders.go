package ordersrepository

import (
	"context"

	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/models"
	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/usecase/orders"
)

func (r *Repository) FindOrders(ctx context.Context, order orders.Filter) ([]*models.Order, error) {
	// switch filter
	return nil, nil
}
