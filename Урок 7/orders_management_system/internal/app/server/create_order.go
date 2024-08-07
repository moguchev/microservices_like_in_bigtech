package server

import (
	"context"

	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/models"
	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/usecase/orders"
	pb "github.com/moguchev/microservices_courcse/orders_management_system/pkg/api/orders_management_system"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// 1. validation
	if err := validateCreateOrderRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 2. convert delivery models to domain models/DTO
	order := newOrderFromPbCreateOrderRequest(req)

	// 3. call usecase
	newOrder, err := s.OrdersUsecase.CreateOrder(ctx, order)
	if err != nil {
		return nil, err // обработается на уровне middleware
	}

	// 4. convert domain models/DTO to delivery models
	response := &pb.CreateOrderResponse{
		OrderId: string(newOrder.ID),
	}

	// 5. return result
	return response, nil
}

func validateCreateOrderRequest(_ *pb.CreateOrderRequest) error {
	//
	return nil
}

func newOrderFromPbCreateOrderRequest(req *pb.CreateOrderRequest) *orders.CreateOrderInfoDTO {
	return &orders.CreateOrderInfoDTO{
		Buyer: models.Buyer{
			ID: models.UserID(req.GetUserId()),
		},
		DeliveryInfo: models.OrderDeliveryInfo{
			DeliveryVariant: models.DeliveryVariant{
				ID: models.DeliveryVariantID(req.GetDeliveryInfo().GetDeliveryVariantId()),
			},
			Date: req.GetDeliveryInfo().GetDeliveryDate().AsTime(),
		},
		Items: newItemsFromPbCreateOrderRequestSKU(req.GetItems()),
	}
}

func newItemsFromPbCreateOrderRequestSKU(items []*pb.CreateOrderRequest_SKU) []models.SKU {
	skus := make([]models.SKU, 0, len(items))
	for _, item := range items {
		skus = append(skus, models.SKU{
			Item: models.Item{
				ID: models.ItemID(item.GetId()),
			},
			Quantity:    uint(item.GetQuantity()),
			WarehouseID: item.GetWarehouseId(),
		})
	}
	return skus
}
