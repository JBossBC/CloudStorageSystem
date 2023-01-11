package handler

import (
	"fileServer/api/internal/logic"
	"fileServer/api/internal/svc"
	"fileServer/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func queryFileInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewQueryFileInfoLogic(r.Context(), svcCtx)
		resp, err := l.QueryFileInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
