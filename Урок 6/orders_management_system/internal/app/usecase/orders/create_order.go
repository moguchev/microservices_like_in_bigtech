package orders

import (
	"context"
	"errors"
	"fmt"

	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/models"
)

var (
	ErrCreateOrderFailed = errors.New("can't create order")
)

func (uc *usecase) CreateOrder(ctx context.Context, orderInfo *CreateOrderInfoDTO) (*models.Order, error) {
	if err := uc.WMS.ReserveSKU(ctx, orderInfo.Items); err != nil {
		return nil, err // здесь мы возвращаем ошибку err
	}

	order := models.NewOrder()
	order.Buyer = orderInfo.Buyer
	order.DeliveryInfo = orderInfo.DeliveryInfo
	order.Items = orderInfo.Items

	if err := uc.OrdersRepository.CreateOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrCreateOrderFailed, err.Error()) // здесь мы возвращаем ошибку ErrCreateOrderFailed
	}

	return order, nil
}
