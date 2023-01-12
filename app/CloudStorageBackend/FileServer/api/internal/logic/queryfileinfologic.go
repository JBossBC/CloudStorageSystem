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

func (l *QueryFileInfoLogic) QueryFileInfo(req *types.QueryReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	resp = types.NewDefaultRes()
	defer func() {
		if err := recover(); err != nil {
			logx.Error(err)
			resp.GetFailedRep("系统出错")
		}
	}()
	owner, ok := req.MetaInfo["creator"].(string)
	if !ok {
		resp.GetFailedRep("请指定用户")
		return resp, err
	}
	files, err := l.svcCtx.FileRpc.QueryFiles(l.ctx, &pb.QueryFileReq{Owner: owner})
	if err != nil {
		logx.Error(err.Error())
		resp.GetFailedRep("查询失败")
		return resp, err
	}
	list, err := util.ConvertRpcFileMetaList(files.List)
	if err != nil {
		logx.Error(err)
		return resp, err
	}
	resp.AddData(list)
	return resp, err
}
