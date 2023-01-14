package main

import (
	"context"
	"fileServer/mq/scheduleClient/internal/config"
	"fileServer/mq/scheduleClient/internal/svc"
	"fileServer/mq/scheduleClient/internal/task"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

var configFile = flag.String("f", "etc/mq-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.DisableStat()
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	server := svc.NewServiceContext(c)
	serContext := task.NewCronScheduler(context.Background(), server)
	serContext.Register()
	if err := server.Scheduler.Run(); err != nil {
		logx.Errorf("!!!MqueueSchedulerErr!!!  run err:%+v", err)
		os.Exit(1)
	}
}
