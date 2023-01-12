package handler

import (
	"net/http"

	"fileServer/api/internal/logic"
	"fileServer/api/internal/svc"
	"fileServer/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func updateFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateFileLogic(r.Context(), svcCtx)
		resp, err := l.UpdateFile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
