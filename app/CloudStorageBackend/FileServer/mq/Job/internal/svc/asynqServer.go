package svc

import (
	"fileServer/mq/Job/internal/config"
	"fmt"
	"github.com/hibiken/asynq"
)

func NewAsynqServer(c config.Config) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			Concurrency: 20,
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task isFailure ========>>>>>> err: %+v \n", err)
				return true
			},
		},
	)
}
