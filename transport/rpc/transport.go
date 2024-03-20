package grpc

import (
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/endpoint"
	pb "github.com/uninus-opensource/uninus-go-grpc-boilerplate/proto"
)

var (
	layoutFormat = "2006-01-02 15:04:05"
)

type (
	grpcUserServer struct {
		getListUser grpctransport.Handler
	}
)

func NewGRPCUserServer(endpoint endpoint.UserEndpoint, logger log.Logger) pb.UserServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcUserServer{
		getListUser: grpctransport.NewServer(
			endpoint.GetListUserEndpoint,
			decodeGetListUserRequest,
			encodeGetListUserResponse,
			append(options, grpctransport.ServerBefore(jwt.GRPCToContext()))...,
		),
	}
}
