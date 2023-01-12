package logic

import (
	"context"

	"fileServer/api/internal/svc"
	"fileServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFileLogic {
	return &DeleteFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFileLogic) DeleteFile(req *types.FindReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	defer func() {
		if panicErr := recover(); panicErr != nil {
			resp.GetFailedRep("系统出错")
		}
	}()
	inputMap := req.MetaInfo
	if _, ok := inputMap["creator"]; !ok {
		resp.GetFailedRep("请指定用户")
		return resp, err
	}
	if _, ok := inputMap["name"]; !ok {
		resp.GetFailedRep("没有要删除的文件")
		return resp, err
	}
	//center logic
	return
}
