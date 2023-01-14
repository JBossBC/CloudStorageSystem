package main

import (
	"context"
	"fileServer/mq/Job/internal/config"
	"fileServer/mq/Job/internal/logic"
	"fileServer/mq/Job/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/asynqserver-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	serverContext := svc.NewServiceContext(c)
	ctx := context.Background()
	cronJob := logic.NewCronJob(ctx, serverContext)
	mux := cronJob.Register()
	if err := serverContext.Scheduler.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
	fmt.Printf("Starting cronJob server at %s:%d...\n", c.Host, c.Port)
}
