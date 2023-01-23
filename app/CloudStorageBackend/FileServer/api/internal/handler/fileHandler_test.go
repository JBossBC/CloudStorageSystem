package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/net"
	"net/http"
	"testing"
)

func BenchmarkDownLoad(t *testing.B) {
	header := http.Header{}
	header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMwMDAwMDAwMTY3NDQ3NTg0OCwiaWF0IjoxNjc0NDc1ODQ4LCJ1c2VyIjoieGl5YW5nIn0.yJ4r9Cxs8OnlbQTuPjvgmat95SuIXvA5nxDMsYfjb3U")
	var JSONObj = make(map[string]interface{})
	JSONObj["metaInfo"] = map[string]interface{}{"filename": "README.md"}
	marshal, _ := json.Marshal(JSONObj)
	body := ioutil.NopCloser(bytes.NewReader(marshal))
	request := &http.Request{
		Method: http.MethodGet,
		URL:    net.FormatURL("http", "localhost", 8889, "/file/api/file/download"),
		Header: header,
		Body:   body,
	}
	client := http.Client{}
	for i := 0; i < t.N; i++ {
		do, err := client.Do(request)
		if err != nil && do != nil {
			return
		}
	}
}
