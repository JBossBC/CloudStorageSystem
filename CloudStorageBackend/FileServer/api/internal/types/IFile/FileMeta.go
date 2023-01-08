package IFile

import (
	"os"
	"time"
)

type FileMeta struct {
	creator     string
	owner       string
	name        string
	create_time time.Time
	update_time time.Time
	isDir       bool
	typeof      string
	authority   string
	size        int64
}

func (fm *FileMeta) GetCreator() string {
	return fm.creator
}
func (fm *FileMeta) GetOwner() string {
	return fm.owner
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
