package IClient

type (
	IDFSClient interface {
		//this map is extraData need be use
		Upload(map[string]interface{}, []byte) (string, error)
		Download(map[string]interface{}) error
	}
)

type DFSClientFactory struct {
}
