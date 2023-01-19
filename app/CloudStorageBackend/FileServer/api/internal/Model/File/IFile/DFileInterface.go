package IFile

import (
	"errors"
	"io"
	"sync"
)

type IFile interface {
	InitFile(uri string, deferCall bool) error
	GetFileData() (io.Reader, error)
}

// File  尽量使用充血模型进行优化，data设为私有变量的原因提供延迟注入数据,同时保证开闭原则 /
type File struct {
	// seem the file descriptor
	MetaInfo *FileMeta
	//support the file protocol
	Uri           string
	Lock          sync.Mutex
	DeferWrapData func() error
	WrapOnce      sync.Once
	DataSource    io.Reader
	IFile
}

func (file *File) isEmpty() bool {
	if file.MetaInfo.size <= 0 {
		return true
	}
	return false
}
func (file *File) GetFileLocation() string {
	return file.Uri
}
func (file *File) InitFile(uri string) error {
	return errors.New("不能使用抽象类")
}

/**
  maybe throws the panic from DeferWrapData func
*/
func (file *File) GetFileData() (data io.Reader, err error) {
	file.WrapOnce.Do(func() {
		err = file.DeferWrapData()
	})
	return file.DataSource, err
}
func (file *File) SetFileData(data io.Reader) {
	file.DataSource = data
}
