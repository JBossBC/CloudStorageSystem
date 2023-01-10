package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/internal/svc"
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneLogic {
	return &FindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOneLogic) FindOne(in *pb.FindFileReq) (*pb.FileMetaInfo, error) {
	// todo: add your logic here and delete this line
	var owner = in.Owner
	var name = in.Name
	if owner == "" || name == "" {
		return nil, sqlc.ErrNotFound
	}
	fileModel, err := l.svcCtx.FileModel.FindOne(l.ctx, owner, name)
	if err != nil {
		return nil, err
	}
	return pb.GetFileMetaInfo(fileModel), nil
}
