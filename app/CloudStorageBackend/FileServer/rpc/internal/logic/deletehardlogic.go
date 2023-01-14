package logic

import (
	"context"
	"time"

	"fileServer/rpc/internal/svc"
	"fileServer/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHardLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHardLogic {
	return &DeleteHardLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteHardLogic) DeleteHard(in *pb.BaseTime) (res *pb.BaseRes, err error) {
	// todo: add your logic here and delete this line
	defer func() {
		if errCatch := recover(); errCatch != nil {
			res.GetFailedRes("系统出错")
			logx.Error(errCatch)
			err = nil
		}
	}()
	res = pb.NewDefaultBaseRes()
	baseTime, err := time.Parse(time.RFC850, in.Date)
	if err != nil {
		logx.Error(err)
		res.GetFailedRes("时间转换出错")
		return res, nil
	}
	l.svcCtx.FileModel.DeleteHard(l.ctx, baseTime)
	return res, nil
}
