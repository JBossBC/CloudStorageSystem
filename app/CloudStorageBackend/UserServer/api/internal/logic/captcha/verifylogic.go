package captcha

import (
	"context"
	"fmt"
	"github.com/dchest/captcha"

	"UserServer/api/internal/svc"
	"UserServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLogic) Verify(req *types.CaptchaVerifyReq) (resp *types.CaptchaVerifyRes, err error) {
	// todo: add your logic here and delete this line
	result := captcha.VerifyString(req.Id, req.Value)
	return &types.CaptchaVerifyRes{
		Result: fmt.Sprint(result),
	}, nil
}
