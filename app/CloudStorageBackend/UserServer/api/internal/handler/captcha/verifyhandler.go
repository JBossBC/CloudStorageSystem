package captcha

import (
	"UserServer/api/internal/logic/captcha"
	"net/http"

	"UserServer/api/internal/svc"
	"UserServer/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func VerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaVerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := captcha.NewVerifyLogic(r.Context(), svcCtx)
		resp, err := l.Verify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
