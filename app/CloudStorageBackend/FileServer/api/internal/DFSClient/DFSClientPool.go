package DFSClient

import (
	"io"
	"sync"
)

// improve memory and defend the file system is break
type fastDFSPool struct {
	schedulerMachine chan *fastDFSClient //number is equal to cacheNums,to avoid the producer waitting
	cacheNums        int                 //default is 100
}
type PoolOption func(pool *fastDFSPool)

var (
	//exposed to the outside world to get fastdfs client
	fastClientpool *fastDFSPool
	once           sync.Once
	lock           sync.Mutex
)

// you should must init the fastClientPool when you use it
func InitFastDFSPool(cacheNums int, fastDFSOption ...FastDFSOption) {
	once.Do(func() {
		fastClientpool = &fastDFSPool{}
		//init pool object for system
		fastClientpool = &fastDFSPool{cacheNums: cacheNums}
		fastClientpool.buildPool(fastDFSOption...)
	})
}

// concurrency is unsafe
func GetFastDFSPool() *fastDFSPool {
	// if you cant use initFastDFSPool to init the fastClientPool before you use GetFastDFSPool function,system considers to help you init this variable by default way
	lock.Lock()
	if fastClientpool == nil {
		InitFastDFSPool(100)
	}
	lock.Unlock()
	return fastClientpool
}

// according to the feature of channel to complete cache queue
func (pool *fastDFSPool) Upload(extraData map[string]interface{}, data []byte) (string, error) {
	client := <-pool.schedulerMachine
	defer func() {
		pool.schedulerMachine <- client
	}()
	URL, err := client.upload(extraData, data)
	if err != nil {
		return "", err
	}
	return URL, nil
}

/**
后期封装为自定义的IFile obj
*/
func (pool *fastDFSPool) Download(uri string) (io.Reader, error) {
	client := <-pool.schedulerMachine
	defer func() {
		pool.schedulerMachine <- client
	}()
	download, err := client.download(uri)
	if err != nil {
		return nil, err
	}
	return download, nil
}

func (dfsPool *fastDFSPool) buildPool(options ...FastDFSOption) {
	dfsPool.schedulerMachine = make(chan *fastDFSClient, dfsPool.cacheNums)
	for i := 0; i < dfsPool.cacheNums; i++ {
		protoClient := newFastDFSClient(options...)
		dfsPool.schedulerMachine <- protoClient
	}
}
