package server

import (
	"context"
	"fmt"
	"github.com/go-kit/log/level"
	ulog "github.com/uninus-opensource/uninus-go-architect-common/log"
	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/model"
	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/utils/constant"
	shv "github.com/uninus-opensource/uninus-go-grpc-boilerplate/utils/sharevar"
)

func (e *Server) GetListUser(ctx context.Context, req model.GetListUserRequest) (model.GetListUserResponse, error) {
	const funcName string = `GetListUser`

	basicLog := genBasicLog(ctx)
	level.Info(shv.Logger).Log(ulog.LogInfo, fmt.Sprintf(`Upper %s`, funcName), ulog.LogReq, fmt.Sprintf(`%+v`, req), ulog.LogBasic, basicLog.Log)

	data, err := e.repo.DBReadWriter.GetListUser(ctx, req)
	if err != nil {
		return model.GetListUserResponse{}, err
	}
	level.Info(shv.Logger).Log(ulog.LogResp, fmt.Sprintf(constant.ExecTime, funcName, ulog.LogTime, basicLog.GetTimeSince()), ulog.LogBasic, basicLog.Log)
	return model.GetListUserResponse{
		Message: data.Message,
		Data:    data.Data,
		Total:   data.Total,
	}, nil
}
