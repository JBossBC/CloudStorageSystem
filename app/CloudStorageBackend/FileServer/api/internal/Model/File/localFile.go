package File

import (
	"bufio"
	"fileServer/api/internal/Model/File/IFile"

	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

type LocalFile struct {
	IFile.File
}

func (lf *LocalFile) InitFile(uri string, deferCall bool) (err error) {
	file, err := os.OpenFile(uri, os.O_APPEND, 644)
	if err != nil {
		return err
	}
	defer func() {
		lf.DataSource = file
		//注册延迟调用函数
		lf.DeferWrapData = func() error {
			lf.Lock.Lock()
			defer lf.Lock.Unlock()
			lf.DataSource = bufio.NewReader(lf.DataSource)
			return nil
		}
		if !deferCall {
			lf.WrapOnce.Do(func() {
				err = lf.DeferWrapData()
			})
		}
		//错误恢复
		if panicErr := recover(); panicErr != nil {
			logx.Errorf("localFile %s init error:%s", uri, panicErr)
		}
	}()
	lf.Uri = uri
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	lf.MetaInfo = IFile.CreateFileMeta(fileInfo)
	return err
}
