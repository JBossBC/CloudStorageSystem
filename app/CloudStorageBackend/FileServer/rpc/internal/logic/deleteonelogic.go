package logic

import (
	"context"
	"time"

	"fileServer/rpc/internal/svc"
	"fileServer/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOneLogic {
	return &DeleteOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOneLogic) DeleteOne(in *pb.FileMetaInfo) (result *pb.BaseRes, err error) {
	// todo: add your logic here and delete this line
	info, err := pb.ConvertFileMetaInfo(in)
	result = pb.NewDefaultBaseRes()
	if err != nil {
		result.GetFailedRes(err.Error())
		return result, err
	}
	info.DeleteTime.Time = time.Now()
	l.svcCtx.FileModel.Update(l.ctx, info)
	return result, nil
}
