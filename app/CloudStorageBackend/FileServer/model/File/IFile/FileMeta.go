package IFile

import (
	"os"
	"time"
)

/**
文件系统中文件的状态
*/
type FileMeta struct {
	name        string
	create_time time.Time
	update_time time.Time
	isDir       bool
	typeof      string
	authority   string
	size        int64
}

func (fm *FileMeta) GetSize() int64 {
	return fm.size
}

func CreateFileMeta(fileInfo os.FileInfo) (metaData *FileMeta) {
	metaData = &FileMeta{}
	metaData.size = fileInfo.Size()
	metaData.update_time = fileInfo.ModTime()
	metaData.isDir = fileInfo.IsDir()
	metaData.name = fileInfo.Name()
	var fileMode = fileInfo.Mode()
	metaData.authority = fileMode.Perm().String()
	metaData.typeof = fileMode.Type().String()
	return metaData
}
