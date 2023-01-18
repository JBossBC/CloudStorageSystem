package File

type IPFSFile struct {
}

func (*IPFSFile) InitFile(uri string, deferCall bool) error {
	return nil
}
func (*IPFSFile) GetFileData() []byte {
	return nil
}
