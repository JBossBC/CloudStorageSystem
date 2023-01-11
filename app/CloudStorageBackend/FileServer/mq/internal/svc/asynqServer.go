package svc

import (
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/mq/internal/config"
	"github.com/hibiken/asynq"
)

func NewAsynqServer(c config.Config) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			Concurrency: 20,
		},
	)
}
