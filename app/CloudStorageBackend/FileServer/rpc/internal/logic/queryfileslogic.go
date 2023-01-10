package logic

import (
	"context"

	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/internal/svc"
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryFilesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFilesLogic {
	return &QueryFilesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryFilesLogic) QueryFiles(in *pb.FindFileReq) (*pb.QueryFileRes, error) {
	// todo: add your logic here and delete this line
    
	return &pb.QueryFileRes{}, nil
}
