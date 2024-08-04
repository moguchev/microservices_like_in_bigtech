package warehousemanagementsystem

import (
	"context"

	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/models"
)

func (c *Client) ReserveSKU(ctx context.Context, items []models.SKU) error {
	return models.ErrNotFound
}
