package go_gcloud_engine

import (
	"encoding/json"
	"net/http"
	"time"
)

func StorageCreateFile(Url, Env string, request *http.Request, Bucket, fileName string, content []byte) {
	Url += "/StorageCreateFile"

	var data1 struct {
		Bucket   string      `json:"bucket"`
		FileName string      `json:"file_name"`
		Content  interface{} `json:"content"`
		TimeNow  string      `json:"time_now"`
	}

	data1.Bucket = Bucket
	data1.FileName = fileName
	_ = json.Unmarshal(content, &data1.Content)
	data1.TimeNow = time.Now().UTC().Format(time.RFC3339Nano)

	var body1, _ = json.Marshal(data1)
	_, _, _, _ = RequestCustomer(Env, "POST", Url, body1, nil, request)

	//if Env == "local" {
	//	return
	//}
	//var ctx = appengine.NewContext(request)
	//var client, err = storage.NewClient(ctx)
	//if err != nil {
	//	return
	//} else {
	//	defer client.Close()
	//}
	//var buf = &bytes.Buffer{}
	//var ws = client.Bucket(Bucket).Object(fileName).NewWriter(ctx)
	//ws.ContentType = "application/json"
	//ws.ContentEncoding = "gzip"
	//var zw = gzip.NewWriter(buf)
	//_, _ = zw.Write(content)
	//_ = zw.Close()
	//_, _ = io.Copy(ws, buf)
	//_ = ws.Close()
}

func StorageCreateFileSVG(Url, Env string, request *http.Request, Bucket, fileName string, content []byte) {
	Url += "/StorageCreateFileSVG"

	var data1 struct {
		Bucket   string      `json:"bucket"`
		FileName string      `json:"file_name"`
		Content  interface{} `json:"content"`
		TimeNow  string      `json:"time_now"`
	}

	data1.Bucket = Bucket
	data1.FileName = fileName
	_ = json.Unmarshal(content, &data1.Content)
	data1.TimeNow = time.Now().UTC().Format(time.RFC3339Nano)

	var body1, _ = json.Marshal(data1)
	_, _, _, _ = RequestCustomer(Env, "POST", Url, body1, nil, request)

	//if Env == "local" {
	//	return
	//}
	//var ctx = appengine.NewContext(request)
	//var client, err = storage.NewClient(ctx)
	//if err != nil {
	//	return
	//} else {
	//	defer client.Close()
	//}
	//var buf = &bytes.Buffer{}
	//var ws = client.Bucket(Bucket).Object(fileName).NewWriter(ctx)
	//ws.ContentType = "image/svg+xml"
	//ws.ContentEncoding = "gzip"
	//var zw = gzip.NewWriter(buf)
	//_, _ = zw.Write(content)
	//_ = zw.Close()
	//_, _ = io.Copy(ws, buf)
	//_ = ws.Close()
}

func StorageCreateMultiFile(Url, Env string, request *http.Request, Bucket string, fList []struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}) {
	Url += "/StorageCreateMultiFile"

	var data1 struct {
		Bucket   string `json:"bucket"`
		FileList []struct {
			Name string `json:"name"`
			Data []byte `json:"data"`
		} `json:"file_list"`
		TimeNow string `json:"time_now"`
	}

	data1.Bucket = Bucket
	data1.FileList = fList
	data1.TimeNow = time.Now().UTC().Format(time.RFC3339Nano)

	var body1, _ = json.Marshal(data1)
	_, _, _, _ = RequestCustomer(Env, "POST", Url, body1, nil, request)

	//if Env == "local" {
	//	return
	//}
	//var ctx = appengine.NewContext(request)
	//var client, err = storage.NewClient(ctx)
	//if err != nil {
	//	return
	//} else {
	//	defer client.Close()
	//}
	//for _, v := range fList {
	//	var buf = &bytes.Buffer{}
	//	var ws = client.Bucket(Bucket).Object(v.Name).NewWriter(ctx)
	//	ws.ContentType = "application/json"
	//	ws.ContentEncoding = "gzip"
	//	var zw = gzip.NewWriter(buf)
	//	_, _ = zw.Write(v.Data)
	//	_ = zw.Close()
	//	_, _ = io.Copy(ws, buf)
	//	_ = ws.Close()
	//}
}

func StorageDeleteFile(Url, Env string, request *http.Request, Bucket string, fileName string) {
	Url += "/StorageDeleteFile"

	var data1 struct {
		Bucket   string `json:"bucket"`
		FileName string `json:"file_name"`
		TimeNow  string `json:"time_now"`
	}

	data1.Bucket = Bucket
	data1.FileName = fileName
	data1.TimeNow = time.Now().UTC().Format(time.RFC3339Nano)

	var body1, _ = json.Marshal(data1)
	_, _, _, _ = RequestCustomer(Env, "POST", Url, body1, nil, request)

	//if Env == "local" {
	//	return
	//}
	//var ctx = appengine.NewContext(request)
	//var client, err = storage.NewClient(ctx)
	//if err != nil {
	//	return
	//} else {
	//	defer client.Close()
	//}
	//_ = client.Bucket(Bucket).Object(fileName).Delete(ctx)
	return
}

func StorageReadFile(Url, Env string, request *http.Request, Bucket string, fileName string) ([]byte, error) {
	Url += "/StorageReadFile"

	var data1 struct {
		Bucket   string `json:"bucket"`
		FileName string `json:"file_name"`
		TimeNow  string `json:"time_now"`
	}

	data1.Bucket = Bucket
	data1.FileName = fileName
	data1.TimeNow = time.Now().UTC().Format(time.RFC3339Nano)

	var body1, _ = json.Marshal(data1)
	_, body2, err2, _ := RequestCustomer(Env, "POST", Url, body1, nil, request)
	return body2, err2

	//if Env == "local" {
	//	return nil, nil
	//}
	//var ctx = appengine.NewContext(request)
	//var client, err = storage.NewClient(ctx)
	//if err != nil {
	//	return nil, err
	//} else {
	//	defer client.Close()
	//}
	//rc, err := client.Bucket(Bucket).Object(fileName).NewReader(ctx)
	//if err != nil {
	//	return nil, err
	//} else {
	//	defer rc.Close()
	//}
	//return ioutil.ReadAll(rc)
}

func StorageCheckFile(Url, Env string, request *http.Request, Bucket string, fileName string) bool {
	Url += "/StorageCheckFile"

	var data1 struct {
		Bucket   string `json:"bucket"`
		FileName string `json:"file_name"`
		TimeNow  string `json:"time_now"`
	}

	data1.Bucket = Bucket
	data1.FileName = fileName
	data1.TimeNow = time.Now().UTC().Format(time.RFC3339Nano)

	var body1, _ = json.Marshal(data1)
	_, body2, err2, _ := RequestCustomer(Env, "POST", Url, body1, nil, request)

	var result struct {
		Success bool `json:"success"`
	}

	if err2 == nil {
		_ = json.Unmarshal(body2, &result)
	}

	return result.Success

	//if Env == "local" {
	//	return false
	//}
	//var ctx = appengine.NewContext(request)
	//var client, err = storage.NewClient(ctx)
	//if err != nil {
	//	return false
	//} else {
	//	defer client.Close()
	//}
	//_, err = client.Bucket(Bucket).Object(fileName).NewReader(ctx)
	//if err != nil {
	//	return false
	//}
	//return true
}

func StorageDeleteMultiFile(Url, Env string, request *http.Request, Bucket string, array1 []string) {
	Url += "/StorageDeleteMultiFile"

	var data1 struct {
		Bucket     string   `json:"bucket"`
		ListDelete []string `json:"list_delete"`
		TimeNow    string   `json:"time_now"`
	}

	data1.Bucket = Bucket
	data1.ListDelete = array1
	data1.TimeNow = time.Now().UTC().Format(time.RFC3339Nano)

	var body1, _ = json.Marshal(data1)
	_, _, _, _ = RequestCustomer(Env, "POST", Url, body1, nil, request)

	//if Env == "local" {
	//	return
	//}
	//var ctx = appengine.NewContext(request)
	//var client, err = storage.NewClient(ctx)
	//if err != nil {
	//	return
	//} else {
	//	defer client.Close()
	//}
	//for _, v := range array1 {
	//	_ = client.Bucket(Bucket).Object(v).Delete(ctx)
	//}
	return
}
