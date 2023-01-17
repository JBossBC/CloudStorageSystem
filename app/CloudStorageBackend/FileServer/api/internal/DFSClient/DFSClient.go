package DFSClient

import (
	"fileServer/api/internal/DFSClient/IClient"
	"github.com/eventials/go-tus"
)

type FastDFSClient struct {
	IClient.IDFSClient
	FastDFSAdapter
}

type FastDFSAdapter struct {
	tus.Client
}
type Option func(client *FastDFSClient)

func NewFastDFSClient(url string, option ...Option) *FastDFSClient {
	return &tus.Client{}
}

func (*FastDFSClient) Upload(map[string]interface{}, byte[]) error {
	return nil
}
func (*FastDFSClient) Download(map[string]interface{}) error {
	return nil
}
