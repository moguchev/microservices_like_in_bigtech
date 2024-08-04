package orders

import "github.com/moguchev/microservices_courcse/orders_management_system/internal/app/models"

type (
	CreateOrderInfoDTO struct {
		Buyer        models.Buyer
		Items        []models.SKU
		DeliveryInfo models.OrderDeliveryInfo
	}

	Filter struct {
		OrderIDs []models.OrderID
		UserID   *models.UserID
		// any other filters
	}
)
