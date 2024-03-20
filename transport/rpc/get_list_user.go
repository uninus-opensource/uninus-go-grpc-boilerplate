package grpc

import (
	"context"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/model"
	pb "github.com/uninus-opensource/uninus-go-grpc-boilerplate/proto"
	"google.golang.org/grpc"
	"time"

	util "github.com/uninus-opensource/uninus-go-architect-common/grcp"
)

func (g *grpcUserServer) GetListUser(ctx context.Context, req *pb.GetListUserRequest) (*pb.GetListUserResponse, error) {
	_, resp, err := g.getListUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetListUserResponse), nil
}

func decodeGetListUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetListUserRequest)
	return model.GetListUserRequest{
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}

func encodeGetListUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(model.GetListUserResponse)
	user := make([]*pb.User, len(resp.Data))
	for i, v := range resp.Data {
		var (
			updatedAt string = ""
			createdAt string = ""
		)
		if v.CreatedAt != 0 {
			createdAt = time.Unix(int64(v.CreatedAt), 0).String()
		}
		if v.UpdatedAt != 0 {
			updatedAt = time.Unix(int64(v.UpdatedAt), 0).String()
		}

		user[i] = &pb.User{
			Id:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
	}

	return &pb.GetListUserResponse{
		Message: resp.Message,
		Data:    user,
		Total:   resp.Total,
	}, nil
}

// client endpoint req & res decode ...
func encodeGetListUserClientRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(model.GetListUserRequest)
	return &pb.GetListUserRequest{
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}

func decodeGetListUserClientResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetListUserResponse)
	user := make([]model.MstUser, len(resp.Data))
	for i, v := range resp.Data {
		createdAt, _ := time.Parse(layoutFormat, v.CreatedAt)
		updatedAt, _ := time.Parse(layoutFormat, v.UpdatedAt)
		user[i] = model.MstUser{
			ID:        v.Id,
			Name:      v.Name,
			Email:     v.Email,
			CreatedAt: int(createdAt.Unix()),
			UpdatedAt: int(updatedAt.Unix()),
		}
	}
	return model.GetListUserResponse{
		Message: resp.Message,
		Data:    user,
		Total:   resp.Total,
	}, nil
}

func makeClientGetLisUserService(
	conn *grpc.ClientConn,
	timeout time.Duration,
	tracer stdopentracing.Tracer,
	logger log.Logger,
) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		getListUserMethodName,
		encodeGetListUserClientRequest,
		decodeGetListUserClientResponse,
		pb.GetListUserResponse{},
		grpctransport.ClientBefore(jwt.ContextToGRPC()),
	).Endpoint()
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(
		util.DefaultCBSetting(getListUserMethodName, timeout),
	))(endpoint)
	return endpoint
}
