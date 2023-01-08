package DFSClient

import "api/internal/DFSClient/IClient"

type FastDFSClient struct {
	IClient.IDFSClient
}

func (*FastDFSClient) Upload(map[string]interface{}, byte[]) error {
	return nil
}
func (*FastDFSClient) Download(map[string]interface{}) error {
	return nil
}
