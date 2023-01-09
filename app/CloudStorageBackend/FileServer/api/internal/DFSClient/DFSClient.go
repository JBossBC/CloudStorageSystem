package DFSClient

import "cloudStorageSystem/app/CloudStorageBackend/FileServer/api/internal/DFSClient/IClient"

type FastDFSClient struct {
	IClient.IDFSClient
}

func (*FastDFSClient) Upload(map[string]interface{}, byte[]) error {
	return nil
}
func (*FastDFSClient) Download(map[string]interface{}) error {
	return nil
}
