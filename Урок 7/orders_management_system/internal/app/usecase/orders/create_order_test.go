package orders

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/models"
	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/usecase/orders/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_usecase_CreateOrder(t *testing.T) {
	var date = time.Now()

	type args struct {
		ctx       context.Context
		orderInfo *CreateOrderInfoDTO
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Order
		wantErr bool
		mock    func(t *testing.T) Deps
	}{
		{
			name: "Test 1. Positive",
			args: args{
				ctx: context.Background(), // dummy
				orderInfo: &CreateOrderInfoDTO{
					Buyer: models.Buyer{ID: 1},
					Items: []models.SKU{
						{
							Item:        models.Item{ID: 10},
							Quantity:    20,
							WarehouseID: 30,
						},
					},
					DeliveryInfo: models.OrderDeliveryInfo{
						DeliveryVariant: models.DeliveryVariant{
							ID: 40,
						},
						Date: date,
					},
				},
			},
			wantErr: false,
			want: &models.Order{
				// ID: 1,
				Buyer: models.Buyer{ID: 1},
				Items: []models.SKU{
					{
						Item:        models.Item{ID: 10},
						Quantity:    20,
						WarehouseID: 30,
					},
				},
				DeliveryInfo: models.OrderDeliveryInfo{
					DeliveryVariant: models.DeliveryVariant{
						ID: 40,
					},
					Date: date,
				},
			},
			mock: func(t *testing.T) Deps {
				wmsMock := mocks.NewWarehouseManagementSystem(t)
				wmsMock.On("ReserveSKU", mock.Anything, []models.SKU{
					{
						Item:        models.Item{ID: 10},
						Quantity:    20,
						WarehouseID: 30,
					},
				}).Return(nil)

				repoMock := NewMockOrdersRepository(t)

				matchOrder := func(o *models.Order) bool {
					return assert.NotNil(t, o) &&
						assert.Equal(t, models.Buyer{ID: 1}, o.Buyer) &&
						assert.Equal(t, []models.SKU{
							{
								Item:        models.Item{ID: 10},
								Quantity:    20,
								WarehouseID: 30,
							},
						}, o.Items) &&
						assert.Equal(t, models.OrderDeliveryInfo{
							DeliveryVariant: models.DeliveryVariant{
								ID: 40,
							},
							Date: date,
						}, o.DeliveryInfo)
				}
				repoMock.On("CreateOrder", mock.Anything, mock.MatchedBy(matchOrder)).Return(nil)

				return Deps{
					WMS:              wmsMock,
					OrdersRepository: repoMock,
				}
			},
		},
		{
			name: "Test 2. Negative. WMS return error",
			args: args{
				ctx:       context.Background(),
				orderInfo: &CreateOrderInfoDTO{},
			},
			want:    nil,
			wantErr: true,
			mock: func(t *testing.T) Deps {
				wmsMock := mocks.NewWarehouseManagementSystem(t)
				wmsMock.On("ReserveSKU", mock.Anything, mock.Anything).
					Return(errors.New("some error")) // stub

				return Deps{
					WMS: wmsMock,
				}
			},
		},
		{
			name: "Test 3. Negative. Repo return error",
			args: args{
				ctx:       context.Background(),
				orderInfo: &CreateOrderInfoDTO{},
			},
			want:    nil,
			wantErr: true,
			mock: func(t *testing.T) Deps {
				wmsMock := mocks.NewWarehouseManagementSystem(t)
				wmsMock.On("ReserveSKU", mock.Anything, mock.Anything).
					Return(nil) // stub

				repoMock := NewMockOrdersRepository(t)
				repoMock.On("CreateOrder", mock.Anything, mock.Anything).Return(errors.New("some error"))
				return Deps{
					WMS:              wmsMock,
					OrdersRepository: repoMock,
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &usecase{
				Deps: tt.mock(t),
			}
			got, err := uc.CreateOrder(tt.args.ctx, tt.args.orderInfo)
			if !tt.wantErr {
				require.NoError(t, err)
			} else {
				require.NotNil(t, err)
			}

			if got != nil {
				got.ID = ""
			}
			require.Equal(t, tt.want, got)
		})
	}
}
