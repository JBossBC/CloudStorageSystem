package captcha

import (
	"UserServer/api/internal/logic/captcha"
	"net/http"

	"UserServer/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := captcha.NewCaptchaLogic(r.Context(), svcCtx)
		err := l.Captcha(w)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
