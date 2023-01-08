package File

import (
	"bufio"
	IFile2 "fileserver/model/File/IFile"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"os"
)

type LocalFile struct {
	IFile2.File
}

func (lf *LocalFile) InitFile(uri string, deferCall bool) error {
	file, err := os.OpenFile(uri, os.O_APPEND, 644)
	if err != nil {
		return err
	}
	defer func() {
		lf.DataSource = file
		//注册延迟调用函数
		lf.DeferWrapData = func() {
			lf.Lock.Lock()
			defer func() {
				lf.Lock.Unlock()
				if deferErr := recover(); deferErr != nil {
					logx.Errorf("读取文件数据失败:%s", deferErr)
				}
			}()
			data, _ := io.ReadAll(bufio.NewReader(lf.DataSource))
			lf.SetFileData(data)
		}
		if !deferCall {
			lf.WrapOnce.Do(lf.DeferWrapData)
		}
		//错误恢复
		if err := recover(); err != nil {
			logx.Errorf("localFile %s init error:%s", uri, err)
		}
	}()
	lf.Uri = uri
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	lf.MetaInfo = IFile2.CreateFileMeta(fileInfo)
	return err

}
