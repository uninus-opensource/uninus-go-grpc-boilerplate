package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/model"
	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/service"
)

func makeGetListUserEndpoint(svc service.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetListUserRequest)
		res, err := svc.GetListUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

func (e UserEndpoint) GetListUser(ctx context.Context, req model.GetListUserRequest) (model.GetListUserResponse, error) {
	resp, err := e.GetListUserEndpoint(ctx, req)
	if err != nil {
		return model.GetListUserResponse{}, err
	}
	response := resp.(model.GetListUserResponse)
	return response, nil
}
