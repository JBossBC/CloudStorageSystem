package logic

import (
	"context"
	"fileServer/mq/Job/internal/svc"
	"fileServer/mq/Job/jobtype"
	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{ctx: ctx, svcCtx: svcCtx}
}

func (l *CronJob) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	mux.Handle(jobtype.ScheduleDeleteMetaInfo, NewFileMetaDeleteHandler(l.svcCtx))
	return mux
}
