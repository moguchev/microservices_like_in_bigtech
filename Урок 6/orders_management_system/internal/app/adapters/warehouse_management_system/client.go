package warehousemanagementsystem

import "github.com/moguchev/microservices_courcse/orders_management_system/internal/app/usecase/orders"

type Client struct {
	// http/grpc/...
}

var (
	_ orders.WarehouseManagementSystem = (*Client)(nil)
)

func NewClient( /**/ ) *Client {
	return &Client{}
}
