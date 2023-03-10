package handler

import (
	"context"
	"errors"
	"fileServer/api/internal/DFSClient"
	"fileServer/api/internal/svc"
	"fileServer/api/internal/types"
	"fileServer/rpc/pb"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
	"time"
)

const MaxTransportByte int64 = 209715200

func uploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp = types.NewDefaultRes()
		defer func() {
			if panicHandler := recover(); panicHandler != nil {
				resp.GetFailedRep("系统错误")
				httpx.OkJson(w, resp)
			}
		}()

		//support max transport size for a request
		err := r.ParseMultipartForm(MaxTransportByte)
		if err != nil {
			resp.GetFailedRep("未识别到文件")
			httpx.OkJson(w, resp)
			return
		}
		form := r.MultipartForm
		err = handlerFiles(svcCtx, form, r)
		if err != nil {
			resp.GetFailedRep(err.Error())
			httpx.OkJson(w, resp)
			return
		}
		//logic  because the r.MultipartForm data is large,consider  to avoid multiple copy ,so reduce the times of convey,and
		// i must do it because the go-zero logic currently  cant support the logic of existing file
		httpx.OkJson(w, resp)

	}
}
func handlerFiles(handler *svc.ServiceContext, form *multipart.Form, r *http.Request) (err error) {
	group := sync.WaitGroup{}
	group.Add(len(form.File))
	errUrl := make(chan string, len(form.File))
	user := r.Context().Value("user").(string)
	for index, _ := range form.File {
		file, m, err := r.FormFile(index)
		//maybe is unsafe
		defer file.Close()
		if err != nil {
			return err
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		// high performance
		go func() {
			defer func() {
				group.Done()
			}()
			extraData := map[string]interface{}{"user": user, "filename": m.Filename}
			upload, err := DFSClient.GetFastDFSPool().Upload(extraData, data)
			if err != nil {
				errUrl <- fmt.Sprintf("filename %s upload failed:%s", m.Filename, err.Error())
				return
			}
			//combine rpc pb.FileMetaInfo
			info := &pb.FileMetaInfo{
				Creator:     user,
				CreateGroup: user,
				Name:        m.Filename,
				CreateTime:  time.Now().Format(time.RFC850),
				Authority:   "644",
				TypeOf:      "file",
				UpdateTime:  time.Now().Format(time.RFC850),
				Size:        m.Size,
				IsDir:       false,
				DeleteTime:  "",
				Description: upload,
			}
			//spin lock,reduce concurrent load using exponential dispersion algorithm
			var retryTimes = 5
			var recoveryTime = 1 * time.Second
			for i := 0; i < retryTimes; i++ {
				ctx, _ := context.WithTimeout(context.Background(), recoveryTime)
				one, err := handler.FileRpc.InertOne(ctx, info)
				if err == nil && one.Result == "true" {
					break
				}
				<-ctx.Done()
				recoveryTime = time.Duration(int64(float64(recoveryTime.Nanoseconds()) * 1.5))
				time.Sleep(recoveryTime)
			}
			// default In retryTimes,the database operation must be successful
		}()
	}
	group.Wait()
	close(errUrl)
	errLength := len(errUrl)
	if errLength != 0 {
		builder := strings.Builder{}
		for i := 0; i < errLength; i++ {
			builder.WriteString(<-errUrl)
		}
		return errors.New(builder.String())
	}
	return nil
}
