package File

import (
	"errors"
	"fileServer/model/File/IFile"

	"strings"
	"sync"
)

const (
	HTTP_PROTOCOL  = "http"
	HTTPS_PROTOCOL = "https"
	FILE_PROTOCOL  = "file"
)

var SupportProtocol = map[string]interface{}{HTTP_PROTOCOL: nil, HTTPS_PROTOCOL: nil, FILE_PROTOCOL: nil}

type fileFactory struct {
}

var (
	fileFactorySingle *fileFactory
	once              sync.Once
)

func GetFileFactory() *fileFactory {
	once.Do(func() {
		if fileFactorySingle == nil {
			fileFactorySingle = &fileFactory{}
		}
	})
	return fileFactorySingle
}

/**
create file model default local file,use the uri create the fileModel.
取数据支持延迟调用
目前只支持本地文件系统还有使用http/https协议的uri
*/
func (*fileFactory) CreateFile(uri string, deferCall bool) (result IFile.IFile, err error) {
	var protocolName = strings.Split(uri, ":")[0]
	if _, ok := SupportProtocol[protocolName]; !ok {
		return nil, errors.New("不支持此类文件协议")
	}
	switch protocolName {
	case "file":
		result = &LocalFile{}
		err = result.InitFile(strings.SplitN(uri, "///", 2)[1], deferCall)
	case "http", "https":
		result = &fastDFile{}
		err = result.InitFile(uri, deferCall)
	default:
		err = errors.New("不能识别的文件协议")
	}
	return result, err
}
func (ff *fileFactory) CreateFileDefer(uri string) (IFile.IFile, error) {
	return ff.CreateFile(uri, true)
}
