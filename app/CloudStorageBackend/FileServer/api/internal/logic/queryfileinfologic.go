package logic

import (
	"context"
	"fileServer/api/internal/svc"
	"fileServer/api/internal/types"
	"fileServer/api/internal/util"
	"fileServer/rpc/pb"

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

func (l *QueryFileInfoLogic) QueryFileInfo(req *types.QueryReq) (resp *types.BaseResponse) {
	// todo: add your logic here and delete this line
	resp = types.NewDefaultRes()
	defer func() {
		if errHandler := recover(); errHandler != nil {
			logx.Error(errHandler)
			resp.GetFailedRep("系统出错")
		}
	}()
	owner, ok := req.MetaInfo["creator"].(string)
	if !ok {
		resp.GetFailedRep("请指定用户")
		return resp
	}
	files, rpcErr := l.svcCtx.FileRpc.QueryFiles(l.ctx, &pb.QueryFileReq{Owner: owner})
	if rpcErr != nil {
		logx.Error(rpcErr.Error())
		resp.GetFailedRep("查询失败")
		return resp
	}
	list, convertErr := util.ConvertRpcFileMetaList(files.List)
	if convertErr != nil {
		logx.Error(convertErr.Error())
		resp.GetFailedRep("系统出错")
		return resp
	}
	resp.AddData(list)
	return resp
}
