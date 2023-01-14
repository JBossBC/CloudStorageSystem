package svc

import (
	"fileServer/mq/scheduleClient/internal/config"
	"fileServer/rpc/fileserver"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	Scheduler     *asynq.Scheduler
	FileserverRpc fileserver.FileServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		FileserverRpc: fileserver.NewFileServer(zrpc.MustNewClient(c.FileserverClient)),
		Scheduler:     NewAsynqClient(c),
	}
}
