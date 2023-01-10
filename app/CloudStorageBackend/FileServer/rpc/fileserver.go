package main

import (
	"flag"
	"fmt"

	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/internal/config"
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/internal/server"
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/internal/svc"
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/fileserver.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterFileServerServer(grpcServer, server.NewFileServerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
