package svc

import (
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/mq/internal/config"
	"fmt"
	"github.com/hibiken/asynq"
)

func NewAsynqServer(c config.Config) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{IsFailure: func(err error) bool {
			fmt.Printf("asynq server exec task is failure>>>>>>>>>>>>>>> err: %+v \n", err)
			return true
		},
			Concurrency: 20,
		},
	)
}
