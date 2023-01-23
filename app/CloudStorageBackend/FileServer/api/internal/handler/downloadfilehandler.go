package handler

import (
	"bufio"
	"fileServer/api/internal/logic"
	"fileServer/api/internal/svc"
	"fileServer/api/internal/types"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func downloadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewDownloadFileLogic(r.Context(), svcCtx)
		resp := l.DownloadFile(&req)
		if resp.Result != "true" {
			httpx.OkJson(w, resp)
			return
		}
		header := w.Header()
		header.Set("Content-Disposition", fmt.Sprintf("attachment;fileName=%s", req.MetaInfo["filename"].(string)))
		reader := resp.Data.(*bufio.Reader)
		all, _ := ioutil.ReadAll(reader)
		w.Write(all)
	}
}
