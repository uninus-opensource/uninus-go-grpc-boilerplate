package service

import (
	"context"

	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/model"
)

type UserServer interface {
	GetListUser(ctx context.Context, params model.GetListUserRequest) (model.GetListUserResponse, error)
}