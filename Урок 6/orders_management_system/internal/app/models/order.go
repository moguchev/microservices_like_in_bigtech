package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderID string

type Order struct {
	ID           OrderID
	Buyer        Buyer
	Items        []SKU
	DeliveryInfo OrderDeliveryInfo
}

func NewOrder() *Order {
	return &Order{
		ID: OrderID(uuid.New().String()),
	}
}

type OrderDeliveryInfo struct {
	DeliveryVariant DeliveryVariant
	Date            time.Time
}
