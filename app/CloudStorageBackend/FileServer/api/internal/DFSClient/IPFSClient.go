package DFSClient

import "github.com/libp2p/go-libp2p/core/host"

type IPFSClient struct {
	host.Host
}

func (*IPFSClient) Download(map[string]interface{}) error {

}

func (*IPFSClient) Upload(map[string]interface{}, byte[]) error {

}
