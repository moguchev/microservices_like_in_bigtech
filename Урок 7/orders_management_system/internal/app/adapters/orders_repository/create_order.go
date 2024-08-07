package ordersrepository

import (
	"context"

	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/models"
)

func (r *Repository) CreateOrder(ctx context.Context, order *models.Order) error {
	return models.ErrAlreadyExists
}
