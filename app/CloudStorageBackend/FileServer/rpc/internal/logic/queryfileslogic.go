package logic

import (
	"context"
	"fileServer/rpc/internal/svc"
	"fileServer/rpc/pb"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

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

func (l *QueryFilesLogic) QueryFiles(in *pb.QueryFileReq) (*pb.QueryFileRes, error) {
	// todo: add your logic here and delete this line
	var owner = in.Owner
	if owner == "" {
		return nil, sqlx.ErrNotFound
	}
	query, err := l.svcCtx.FileModel.Query(l.ctx, owner)
	if err != nil {
		return nil, err
	}
	return &pb.QueryFileRes{
		Result: &pb.BaseRes{Result: "true", Message: "处理成功"},
		List:   pb.GetFileMetaList(query),
	}, nil
}
