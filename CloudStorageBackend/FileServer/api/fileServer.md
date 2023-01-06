 ### 1. "uploadFile"

1. route definition

- Url: /file/api/file/upload
- Method: POST
- Request: `uploadReq`
- Response: `baseResponse`

2. request definition



```golang
type UploadReq struct {
	MetaInfo map[string]interface{} `json:"metaInfo"`
	Data []byte `form:"fileData"`
}
```


3. response definition



```golang
type BaseResponse struct {
	Result string `json:"result"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}
```

### 2. "downloadFile"

1. route definition

- Url: /file/api/file/download
- Method: GET
- Request: `downloadReq`
- Response: `baseResponse`

2. request definition



```golang
type DownloadReq struct {
	MetaInfo map[string]interface{} `json:"metaInfo"` //must include the location key and the
}
```


3. response definition



```golang
type BaseResponse struct {
	Result string `json:"result"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}
```

### 3. "getFileInfo"

1. route definition

- Url: /file/api/file/get
- Method: GET
- Request: `findReq`
- Response: `baseResponse`

2. request definition



```golang
type FindReq struct {
	MetaInfo map[string]interface{} `json:"metaInfo"` //must include the location key
}
```


3. response definition



```golang
type BaseResponse struct {
	Result string `json:"result"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}
```

### 4. "queryFileInfo"

1. route definition

- Url: /file/api/file/query
- Method: GET
- Request: `queryReq`
- Response: `baseResponse`

2. request definition



```golang
type QueryReq struct {
	MetaInfo map[string]interface{} `json:"metaInfo"` //mush include the owner name
}
```


3. response definition



```golang
type BaseResponse struct {
	Result string `json:"result"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}
```

