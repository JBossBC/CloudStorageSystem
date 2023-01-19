package DFSClient

import (
	"fmt"
	"testing"
)

func TestTransportFile(t *testing.T) {
	url, err := GetFastDFSPool().Upload(map[string]interface{}{"filename": "big", "user": "xiyang"}, []byte("jiangxiyanssssg"))
	if err != nil {
		print(err.Error())
		return
	}
	fmt.Println(url)
}
