package File

import (
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/model/File/IFile"
	"errors"
)

type fastDFile struct {
	IFile.File
}

func (df *fastDFile) InitFile(uri string, deferCall bool) error {
	return errors.New("目前不支持 fastdfs 协议")
}
