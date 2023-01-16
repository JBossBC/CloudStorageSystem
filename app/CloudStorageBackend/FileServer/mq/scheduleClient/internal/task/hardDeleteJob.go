package task

import (
	"fileServer/mq/Job/jobtype"
	"github.com/hibiken/asynq"
)

func (l *MqueueScheduler) hardDeleteFileMeta() {
	task := asynq.NewTask(jobtype.ScheduleDeleteMetaInfo, nil)
	//重试两次，避免排他锁造成性能下降的很快,这里避免使用指数退避算法，考虑到删除这个可能会成为一个大体量的操作，会造成 其他服务的服务停止从而影响系统可用性
	l.svcCtx.Scheduler.Register("*/1 * * * *", task, asynq.MaxRetry(2), asynq.Queue("FileDeleteSchedule"))
}
