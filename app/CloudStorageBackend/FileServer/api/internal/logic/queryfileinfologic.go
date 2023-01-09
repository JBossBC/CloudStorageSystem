package logic

import (
	"context"
	"fileserver/api/internal/svc"
	"fileserver/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryFileInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFileInfoLogic {
	return &QueryFileInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryFileInfoLogic) QueryFileInfo(req *types.QueryReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
