package DFSClient

import (
	"io/ioutil"
	"sync"
	"testing"
)

import _ "net/http/pprof"

func TestUploadConcurrency(t *testing.T) {
	const MaxTestTimes = 10000000
	group := sync.WaitGroup{}
	for i := 0; i < MaxTestTimes; i++ {
		go func() {
			group.Add(1)
			defer func() {
				group.Done()
			}()
			GetFastDFSPool().Upload(nil, nil)
		}()

	}
	group.Wait()

}

func BenchmarkConcurrencyUpload(b *testing.B) {
	group := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		go func() {
			group.Add(1)
			defer func() {
				group.Done()
			}()
			GetFastDFSPool().Upload(nil, nil)
		}()
	}
	group.Wait()
}
func BenchmarkUpload(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetFastDFSPool().Upload(nil, nil)
	}
}
func BenchmarkDownload(t *testing.B) {
	t.N = 50000
	group := sync.WaitGroup{}
	group.Add(t.N)
	for i := 0; i < t.N; i++ {
		go func() {
			defer func() {
				group.Done()
			}()
			download, err := GetFastDFSPool().Download("http://127.0.0.1:8080/group1/big/upload/9bb7f4ea7472846e226a716c35ca8fb4")
			if err != nil {
				print(err.Error())
				return
			}
			_, err = ioutil.ReadAll(download)
			if err != nil {
				print(err.Error())
				return
			}
		}()
	}
}

func BenchmarkDownloadByRoutinue(t *testing.B) {
	t.N = 50000
	group := sync.WaitGroup{}
	group.Add(t.N)
	for i := 0; i < t.N; i++ {
		go func() {
			defer func() {
				group.Done()
			}()
			download, err := newFastDFSClient().download("http://127.0.0.1:8080/group1/big/upload/d825b677594299b6f15cd58ebeb09530")
			if err != nil {
				//print(err.Error())
				return
			}
			_, err = ioutil.ReadAll(download)
			if err != nil {
				//print(err.Error())
				return
			}
		}()

	}
	group.Wait()
}
