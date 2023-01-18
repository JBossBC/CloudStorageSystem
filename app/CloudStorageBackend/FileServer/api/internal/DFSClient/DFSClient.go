package DFSClient

import (
	"errors"
	"github.com/eventials/go-tus"
	"os"
	"strings"
)

var FastDFSCluster string = "http://192.168.1.10:8080/"

func init() {
	getenv := os.Getenv("default-DFSCluster-Location")
	if getenv == "" && FastDFSCluster == "" {
		panic("Cant find the  fastDFS cluster  location,running fileserver failed")
	}
	if FastDFSCluster == "" {
		FastDFSCluster = getenv
	}
}

type (
	FastDFSClient struct {
		Client *tus.Client
	}
	FastDFSOption func(client *FastDFSClient)
)

/**
  url set nil,beacuse the url decide the upload file location,需要单例模式
*/
func NewFastDFSClient(option ...FastDFSOption) (result *FastDFSClient) {
	client, _ := tus.NewClient("", nil)
	if !strings.HasSuffix(FastDFSCluster, "/") {
		FastDFSCluster = FastDFSCluster + "/"
	}
	result = &FastDFSClient{
		Client: client,
	}
	for i, _ := range option {
		option[i](result)
	}
	return result
}

func (FC *FastDFSClient) Upload(extraData map[string]interface{}, data []byte) error {
	filename, ok := extraData["filename"].(string)
	if !ok {
		return errors.New("cant find the filename")
	}
	user, ok := extraData["user"].(string)
	if !ok {
		return errors.New("cant find the user ")
	}
	restURL := FC.buildUploadURL(user)
	//set restURL to transport
	FC.Client.Url = restURL
	defer func() {
		FC.Client.Url = FastDFSCluster
	}()
	upload, err := FC.convertFileUpload(filename, data)
	if err != nil {
		return err
	}
	err = upload.Upload()
	if err != nil {
		return err
	}
	return err
}
func (FC *FastDFSClient) convertFileUpload(targetFilename string, data []byte) (result *tus.Uploader, err error) {
	upload := tus.NewUploadFromBytes(data)
	upload.Metadata["filename"] = targetFilename
	result, err = FC.Client.CreateOrResumeUpload(upload)
	return result, err
}
func (FC *FastDFSClient) Download(map[string]interface{}) error {

	return nil
}

func (FC *FastDFSClient) buildUploadURL(creator string) string {
	builder := strings.Builder{}
	builder.WriteString(FastDFSCluster)
	builder.WriteString(creator)
	builder.WriteString("/upload/")
	return builder.String()
}
func WithChunkSize(size int64) FastDFSOption {
	return func(client *FastDFSClient) {
		client.Client.Config.ChunkSize = size
	}
}
func WithResume(flag bool) FastDFSOption {
	return func(client *FastDFSClient) {
		client.Client.Config.Resume = flag
	}
}
func WithOverridePatchMethod(flag bool) FastDFSOption {
	return func(client *FastDFSClient) {
		client.Client.Config.OverridePatchMethod = flag

	}
}
func WithStore(store tus.Store) FastDFSOption {
	return func(client *FastDFSClient) {
		client.Client.Config.Store = store
	}
}
func WithHttpHeader(header map[string]interface{}) FastDFSOption {
	return func(client *FastDFSClient) {
		for key, value := range header {
			client.Client.Config.Header.Set(key, value.(string))
		}
	}
}
