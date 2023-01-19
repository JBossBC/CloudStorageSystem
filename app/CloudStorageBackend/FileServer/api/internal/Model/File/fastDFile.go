package File

import (
	"fileServer/api/internal/DFSClient"
	"fileServer/api/internal/Model/File/IFile"
)

type fastDFile struct {
	IFile.File
}

func (df *fastDFile) InitFile(uri string, deferCall bool) (err error) {
	df.DeferWrapData = func() error {
		download, wrapEerr := DFSClient.GetFastDFSPool().Download(uri)
		if wrapEerr != nil {
			return wrapEerr
		}
		df.DataSource = download
		return nil
	}
	if !deferCall {
		err = df.DeferWrapData()
	}
	return nil
}
