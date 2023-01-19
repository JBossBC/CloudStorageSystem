package DFSClient

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestTransportFile(t *testing.T) {
	url, err := GetFastDFSPool().Upload(map[string]interface{}{"filename": "ss", "user": "xiyang"}, []byte("jiangxiyanssssg"))
	if err != nil {
		print(err.Error())
		return
	}
	fmt.Println(url)
}

func TestFastDFSPool_Download(t *testing.T) {
	download, err := GetFastDFSPool().Download("http://127.0.0.1:8080/group1/big/upload/d825b677594299b6f15cd58ebeb09530")
	if err != nil {
		print(err.Error())
		return
	}
	all, err := ioutil.ReadAll(download)
	if err != nil {
		print(err.Error())
		return
	}
	print(string(all))
}
func BenchmarkDownloadClient(t *testing.B) {
	t.N = 5000
	for i := 0; i < t.N; i++ {
		download, err := newFastDFSClient().download("http://127.0.0.1:8080/group1/big/upload/d825b677594299b6f15cd58ebeb09530")
		if err != nil {
			print(err.Error())
			return
		}
		_, err = ioutil.ReadAll(download)
		if err != nil {
			print(err.Error())
			return
		}
	}
}
