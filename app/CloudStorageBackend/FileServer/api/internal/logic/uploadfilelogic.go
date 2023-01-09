package logic

import (
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/api/internal/svc"
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(req *types.UploadReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	return
}
