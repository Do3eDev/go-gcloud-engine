package go_gcloud_engine

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func StorageCreateFile(
	Url, Env string,
	request *http.Request,
	Bucket, fileName string,
	content []byte,
) {
	var thisTmp = "/var/www/" + Bucket
	var folderName = filepath.Dir("/" + fileName)
	if _, err := os.Stat(thisTmp); os.IsNotExist(err) {
		_ = os.Mkdir(thisTmp, 0755)
	}
	if _, err := os.Stat(thisTmp + folderName); os.IsNotExist(err) {
		_ = os.Mkdir(thisTmp+folderName, 0755)
	}
	reSTT, errTT := os.Create(thisTmp + "/" + fileName)
	if errTT == nil {
		var buf = &bytes.Buffer{}
		var zw = gzip.NewWriter(buf)
		_, _ = zw.Write(content)
		_ = zw.Close()
		_, _ = io.Copy(reSTT, buf)
	}

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
}

func StorageCreateFileSVG(
	Url, Env string,
	request *http.Request,
	Bucket, fileName string,
	content []byte,
) {
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
}

func StorageCreateMultiFile(Url, Env string, request *http.Request, Bucket string, fList []struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}) {
	Url += "/StorageCreateMultiFile"

	var data1 struct {
		Bucket   string `json:"bucket"`
		FileList []struct {
			Name string      `json:"name"`
			Data interface{} `json:"data"`
		} `json:"file_list"`
		TimeNow string `json:"time_now"`
	}

	data1.Bucket = Bucket
	data1.FileList = fList
	data1.TimeNow = time.Now().UTC().Format(time.RFC3339Nano)

	var body1, _ = json.Marshal(data1)
	_, _, _, _ = RequestCustomer(Env, "POST", Url, body1, nil, request)
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
	return
}

func StorageReadFile(
	Url, Env string,
	request *http.Request,
	Bucket string,
	fileName string,
) ([]byte, error) {
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
}

func StorageDeleteMultiFile(
	Url, Env string,
	request *http.Request,
	Bucket string,
	array1 []string,
) {
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
	return
}
