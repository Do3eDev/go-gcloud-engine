# go-gcloud-engine

```go
package main

import (
	"encoding/json"
	"github.com/Do3eDev/go-gcloud-engine"
	"google.golang.org/appengine"
	"net/http"
)

//var Env = "local"
//var Env = "production"
var Env = "staging"

func main() {
	http.HandleFunc("/", test)
	appengine.Main()
}

func test(writer http.ResponseWriter, request *http.Request) {
	add := go_gcloud_engine.MemCacheAdd(Env, request, "key_cache_123", []byte("value_000"), 5) // 5 minutes
	//add := go_gcloud_engine.MemCacheAddSecond(Env, request, "key_cache_123", []byte("value_000"), 15) // 15 seconds
	sb1, _ := json.Marshal(add)
	_, _ = writer.Write(sb1)

	// https://storage.googleapis.com/asset-staging-demo/foler-demo/file-name-demo.txt
	var bucket = "asset-staging-demo"
	var fileName = "foler-demo/file-name-demo.txt"
	var content = "Hello storage"
	go_gcloud_engine.StorageCreateFile(Env, request, bucket, fileName, []byte(content))

	var status, rspBody, err, rspHeader = go_gcloud_engine.RequestCustomer(Env, "POST", "https://test-request.domain.demo/api-path?abc=test&dev=true",
		[]byte(`{"data":"json","test":true}`), map[string]string{"Content-Type": "application/json"}, request)

	if status == 200 && err == nil {
		_, _ = writer.Write(rspBody)

		sb2, _ := json.Marshal(rspHeader)
		_, _ = writer.Write(sb2)
	} else {
		_, _ = writer.Write([]byte(err.Error()))

		go_gcloud_engine.WriteLogError(Env, request, "Test write error: %v", err)
	}
}
```