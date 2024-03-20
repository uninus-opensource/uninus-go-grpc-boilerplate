package server

import (
	"context"
	ulog "github.com/uninus-opensource/uninus-go-architect-common/log"
	mvc "github.com/uninus-opensource/uninus-go-architect-common/microservice"
	repo "github.com/uninus-opensource/uninus-go-grpc-boilerplate/repository"
	sv "github.com/uninus-opensource/uninus-go-grpc-boilerplate/service"
	"time"
)

type Server struct {
	repo repo.Repo
}

func NewUserServer(repo repo.Repo) sv.UserServer {
	return &Server{
		repo: repo,
	}
}

func genBasicLog(ctx context.Context) (basicLog ulog.ConsoleLog) {
	basicLog = ulog.ConsoleLog{
		RequestID: mvc.GetRequestIDByContext(ctx),
		TimeStart: time.Now(),
	}
	basicLog.GenerateConsoleLog(ctx)
	return
}
