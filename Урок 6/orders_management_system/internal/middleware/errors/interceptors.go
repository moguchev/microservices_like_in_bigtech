package middleware

import (
	"context"
	"errors"

	"github.com/moguchev/microservices_courcse/orders_management_system/internal/app/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorsUnaryInterceptor - convert any arror to rpc error
func ErrorsUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			if _, ok := status.FromError(err); ok {
				return resp, err
			}

			switch {
			case errors.Is(models.ErrNotFound, err):
				return nil, status.Error(codes.NotFound, err.Error())
			case errors.Is(models.ErrAlreadyExists, err):
				return nil, status.Error(codes.AlreadyExists, err.Error())
				//
			default:
				return nil, status.Error(codes.Internal, err.Error())
			}

		}

		return
	}
}
