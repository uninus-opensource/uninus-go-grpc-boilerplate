package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/service"
)


type UserEndpoint struct {
	GetListUserEndpoint endpoint.Endpoint
}

func NewUserEndpoint(svc service.UserServer) (UserEndpoint, error) {
	var GetListUserEndpoint endpoint.Endpoint
	{
		GetListUserEndpoint = makeGetListUserEndpoint(svc)
	}

	return UserEndpoint{
		GetListUserEndpoint: GetListUserEndpoint,
	}, nil
}