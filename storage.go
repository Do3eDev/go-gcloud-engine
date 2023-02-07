package go_gcloud_engine

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func StorageCreateFile(
	Url, Env string,
	request *http.Request,
	Bucket, fileName string,
	content []byte,
) {
	var thisTmp = "/var/www/autoketing-storage/" + Bucket
	var folderName = filepath.Dir("/" + fileName)
	var thisPath = ""
	for _, folderName := range strings.Split(thisTmp+folderName, "/") {
		if folderName != "" {
			thisPath += "/" + folderName
			if _, err := os.Stat(thisPath); os.IsNotExist(err) {
				_ = os.Mkdir(thisPath, 0755)
			}
		}
	}
	var filePath = thisTmp + "/" + fileName
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_WRONLY, 0755)
	if err == nil {
		_, _ = f.Write(content)
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
