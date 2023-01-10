package main

import (
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/mq/internal/config"
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/mq/internal/svc"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/mq-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := svc.NewAsynqServer(c)
}
