package util

import (
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/pb"
	"fmt"
	"testing"
	"time"
)

func TestConvertRpcFileMeta(t *testing.T) {
	meta, err := ConvertRpcFileMeta(&pb.FileMetaInfo{
		Creator:     "xiyang",
		CreateGroup: "xiyang",
		Name:        "hello",
		CreateTime:  time.Now().String(),
		Authority:   "644",
		TypeOf:      "file",
		UpdateTime:  time.Now().String(),
		Size:        0,
		IsDir:       false,
		DeleteTime:  time.Now().String(),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for index, value := range meta {
		fmt.Printf("%s %v\n", index, value)
	}
}
