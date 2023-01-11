package handler

import (
	"fileServer/api/internal/logic"
	"fileServer/api/internal/svc"
	"fileServer/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func uploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
