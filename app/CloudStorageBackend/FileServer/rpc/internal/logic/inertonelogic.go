package logic

import (
	"context"
	"fileServer/rpc/internal/svc"
	"fileServer/rpc/pb"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type InertOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInertOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InertOneLogic {
	return &InertOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InertOneLogic) InertOne(in *pb.FileMetaInfo) (result *pb.BaseRes, err error) {
	// todo: add your logic here and delete this line
	info, err := pb.ConvertFileMetaInfo(in)
	result = pb.NewDefaultBaseRes()
	if err != nil {
		result.GetFailedRes(err.Error())
		return result, err
	}
	insert, err := l.svcCtx.FileModel.Insert(l.ctx, info)
	if row, _ := insert.RowsAffected(); err != nil || row <= 0 {
		result.GetFailedRes(fmt.Sprintf("fileRPCServer insertOne error:%v,影响行数:%d", err, row))
		return result, err
	}
	return result, err
}
