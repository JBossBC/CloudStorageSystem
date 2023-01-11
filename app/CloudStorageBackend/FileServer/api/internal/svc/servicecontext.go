package svc

import (
	"fileServer/api/internal/config"
	"fileServer/rpc/fileserver"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	FileRpc fileserver.FileServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		FileRpc: fileserver.NewFileServer(zrpc.MustNewClient(c.FileServerRpcConfig)),
	}
}
