package models

type (
	ItemID      uint64
	WarehouseID uint64
)

type Item struct {
	ID ItemID
}

type SKU struct {
	Item
	Quantity    uint
	WarehouseID uint64
}
