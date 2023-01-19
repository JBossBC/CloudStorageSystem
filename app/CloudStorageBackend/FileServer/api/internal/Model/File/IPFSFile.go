package File

import "io"

type IPFSFile struct {
}

func (*IPFSFile) InitFile(uri string, deferCall bool) error {
	return nil
}
func (*IPFSFile) GetFileData() (io.Reader, error) {
	return nil, nil
}
