package logic

import (
	"context"
	"fileServer/api/internal/DFSClient"
	"fileServer/api/internal/svc"
	"fileServer/api/internal/types"
	"fileServer/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadFileLogic {
	return &DownloadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadFileLogic) DownloadFile(req *types.DownloadReq) (resp *types.BaseResponse) {
	// todo: add your logic here and delete this line
	resp = types.NewDefaultRes()
	defer func() {
		if panicErr := recover(); panicErr != nil {
			logx.Error(panicErr)
			resp.GetFailedRep("系统错误")
		}
	}()
	user, ok := l.ctx.Value("user").(string)
	filename, ok := req.MetaInfo["filename"].(string)
	if !ok {
		resp.GetFailedRep("未指定文件名")
		return resp
	}
	one, err := l.svcCtx.FileRpc.FindOne(l.ctx, &pb.FindFileReq{Owner: user, Name: filename})
	if err != nil {
		resp.GetFailedRep("查询失败")
		return resp
	}
	if one == nil {
		resp.GetFailedRep("文件不存在")
		return resp
	}
	download, err := DFSClient.GetFastDFSPool().Download(one.Description)
	if err != nil {
		resp.GetFailedRep("文件下载失败")
		return resp
	}
	resp.AddData(download)
	return resp
}
