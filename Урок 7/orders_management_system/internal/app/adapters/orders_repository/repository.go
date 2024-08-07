package ordersrepository

import "github.com/moguchev/microservices_courcse/orders_management_system/internal/app/usecase/orders"

type Repository struct {
	// postgres/mongodb/redis/...
}

var (
	_ orders.OrdersRepository = (*Repository)(nil)
)

func NewRepository( /**/ ) *Repository {
	return &Repository{}
}
