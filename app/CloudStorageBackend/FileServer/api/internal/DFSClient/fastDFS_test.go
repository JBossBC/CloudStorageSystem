package DFSClient

import (
	"fmt"
	"testing"
)

func TestTransportFile(t *testing.T) {
	url, err := GetFastDFSPool().Upload(map[string]interface{}{"filename": "hello", "user": "group1"}, []byte("jiangxiyang"))
	if err != nil {
		print(err.Error())
		return
	}
	fmt.Println(url)
}
