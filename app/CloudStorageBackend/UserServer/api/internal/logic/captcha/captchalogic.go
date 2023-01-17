package captcha

import (
	"context"
	"github.com/dchest/captcha"
	"net/http"

	"UserServer/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha(w http.ResponseWriter) error {
	// todo: add your logic here and delete this line
	// 生成验证码id
	id := captcha.NewLen(6)
	w.Header().Add("id", id)
	// 生成验证码并写入`w`
	err := captcha.WriteImage(w, id, 120, 50)
	return err
}
