package task

import (
	"context"
	"fileServer/mq/scheduleClient/internal/svc"
)

type MqueueScheduler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronScheduler(ctx context.Context, svcCtx *svc.ServiceContext) *MqueueScheduler {
	return &MqueueScheduler{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MqueueScheduler) Register() {
	l.hardDeleteFileMeta()
}
