package DFSClient

import (
	"sync"
	"testing"
)

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
	//group.Wait()

}
