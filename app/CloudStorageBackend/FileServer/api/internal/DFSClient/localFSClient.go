package DFSClient

import "fileServer/api/internal/DFSClient/IClient"

type LocalFSClient struct {
	IClient.IDFSClient
}

func (*LocalFSClient) Download(metaInfo map[string]interface{}) error {
	return nil
}
func (*LocalFSClient) Upload(map[string]interface{}, []byte) error {
	return nil
}
