package logic

import (
	"context"
	"fileServer/api/internal/svc"
	"fileServer/api/internal/types"
	"fileServer/api/internal/util"
	"fileServer/rpc/pb"

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
	resp = types.NewDefaultRes()
	defer func() {
		if errHandler := recover(); errHandler != nil {
			logx.Error(errHandler)
			resp.GetFailedRep("参数类型不匹配")
		}
	}()
	owner, ok := req.MetaInfo["creator"].(string)
	if !ok {
		resp.GetFailedRep("未指定用户")
		return resp, err
	}
	name, ok := req.MetaInfo["name"].(string)
	if !ok {
		resp.GetFailedRep("没有指定文件")
		return resp, err
	}
	result, err := l.svcCtx.FileRpc.FindOne(l.ctx, &pb.FindFileReq{Owner: owner, Name: name})
	if err != nil {
		resp.GetFailedRep("获取文件失败")
		return resp, err
	}
	meta, err := util.ConvertRpcFileMeta(result)
	if err != nil {
		logx.Error(err.Error())
		resp.GetFailedRep("系统出错")
		return resp, err
	}
	resp.AddData(meta)
	return resp, err
}
