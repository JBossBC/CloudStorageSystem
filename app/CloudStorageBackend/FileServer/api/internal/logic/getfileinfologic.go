package logic

import (
	"context"
	"fileserver/api/internal/svc"
	"fileserver/api/internal/types"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileInfoLogic {
	return &GetFileInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileInfoLogic) GetFileInfo(req *types.FindReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("-----------------------------------HELLO----------------------------------------")
	return
}
