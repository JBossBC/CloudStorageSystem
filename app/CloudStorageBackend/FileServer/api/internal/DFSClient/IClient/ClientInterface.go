package IClient

type IDFSClient interface {
	Upload(map[string]interface{}, byte[]) error
	Download(map[string]interface{}) error
}

type DFSClientFactory struct {
}
