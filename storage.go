package go_gcloud_engine

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

func StorageCreateFile(
	Url, Env string,
	request *http.Request,
	Bucket, fileName string,
	content []byte,
) {
	var thisTmp = "/var/www/autoketing-storage/" + Bucket
	var filePath = thisTmp + "/" + fileName
	_ = os.MkdirAll(filepath.Dir(filePath), 0o755)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
		if err == nil {
			_, _ = f.Write(content)
		}
	} else {
		_ = os.WriteFile(filePath, content, 0644)
	}
}

func StorageCreateFileSVG(
	Url, Env string,
	request *http.Request,
	Bucket, fileName string,
	content []byte,
) {
	var thisTmp = "/var/www/autoketing-storage/" + Bucket
	var filePath = thisTmp + "/" + fileName
	_ = os.MkdirAll(filepath.Dir(filePath), 0o755)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
		if err == nil {
			_, _ = f.Write(content)
		}
	} else {
		_ = os.WriteFile(filePath, content, 0644)
	}
}

func StorageCreateMultiFile(Url, Env string, request *http.Request, Bucket string, fList []struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}) {
	for _, s := range fList {
		content, _ := json.Marshal(s.Data)
		fileName := s.Name
		var thisTmp = "/var/www/autoketing-storage/" + Bucket
		var filePath = thisTmp + "/" + fileName
		_ = os.MkdirAll(filepath.Dir(filePath), 0o755)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
			if err == nil {
				_, _ = f.Write(content)
			}
		} else {
			_ = os.WriteFile(filePath, content, 0644)
		}
	}
}

func StorageDeleteFile(Url, Env string, request *http.Request, Bucket string, fileName string) {
	return
}

func StorageReadFile(
	Url, Env string,
	request *http.Request,
	Bucket string,
	fileName string,
) ([]byte, error) {
	return nil, nil
}

func StorageCheckFile(Url, Env string, request *http.Request, Bucket string, fileName string) bool {
	var checked bool
	var thisTmp = "/var/www/autoketing-storage/" + Bucket
	var filePath = thisTmp + "/" + fileName
	_ = os.MkdirAll(filepath.Dir(filePath), 0o755)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		checked = true
	}
	return checked
}

func StorageDeleteMultiFile(
	Url, Env string,
	request *http.Request,
	Bucket string,
	array1 []string,
) {
	return
}
