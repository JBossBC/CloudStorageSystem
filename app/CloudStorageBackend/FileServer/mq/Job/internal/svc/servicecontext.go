package svc

import (
	"fileServer/mq/Job/internal/config"
	"fileServer/rpc/fileserver"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	Scheduler     *asynq.Server
	FileserverRpc fileserver.FileServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		FileserverRpc: fileserver.NewFileServer(zrpc.MustNewClient(c.FileserverClient)),
		Scheduler:     NewAsynqServer(c),
	}
}
