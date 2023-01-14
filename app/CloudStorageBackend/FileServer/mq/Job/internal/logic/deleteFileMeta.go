package logic

import (
	"context"
	"errors"
	"fileServer/mq/Job/internal/svc"
	"fileServer/rpc/pb"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type FileMetaDeleteHandler struct {
	svcCtx *svc.ServiceContext
}

func NewFileMetaDeleteHandler(ctx *svc.ServiceContext) *FileMetaDeleteHandler {
	return &FileMetaDeleteHandler{svcCtx: ctx}
}
func (l *FileMetaDeleteHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	expireTime := time.UnixMilli(time.Now().UnixMilli() - time.Hour.Milliseconds())
	hard, err := l.svcCtx.FileserverRpc.DeleteHard(ctx, &pb.BaseTime{Date: expireTime.String()})
	if err != nil || hard.Result != "true" {
		logx.Errorf("定时删除filemeta错误:%v", err)
		return errors.New("delete task failed,maybe should retry")
	}
	return nil
}
