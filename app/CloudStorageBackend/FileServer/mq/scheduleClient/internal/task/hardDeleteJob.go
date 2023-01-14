package task

import (
	"fileServer/mq/Job/jobtype"
	"github.com/hibiken/asynq"
)

func (l *MqueueScheduler) hardDeleteFileMeta() {
	task := asynq.NewTask(jobtype.ScheduleDeleteMetaInfo, nil)
	l.svcCtx.Scheduler.Register("* * */1 * *", task)
}
