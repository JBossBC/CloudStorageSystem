package DFSClient

import "testing"

func TestTransportFile(t *testing.T) {
	println(NewFastDFSClient().Upload(map[string]interface{}{"filename": "hello", "user": "group"}, []byte("jiangxiyang")))
}
