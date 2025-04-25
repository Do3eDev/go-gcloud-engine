package go_gcloud_engine

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

func StorageCreateFile(
	_, _ string,
	_ *http.Request,
	Bucket, fileName string,
	content []byte,
) {
	thisTmp := "/var/www/autoketing-storage/" + Bucket
	filePath := thisTmp + "/" + fileName
	_ = os.MkdirAll(filepath.Dir(filePath), 0o755)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0o644)
		if err == nil {
			_, _ = f.Write(content)
		}
	} else {
		_ = os.WriteFile(filePath, content, 0o644)
	}
}

func StorageCreateFileSVG(
	_, _ string,
	_ *http.Request,
	Bucket, fileName string,
	content []byte,
) {
	thisTmp := "/var/www/autoketing-storage/" + Bucket
	filePath := thisTmp + "/" + fileName
	_ = os.MkdirAll(filepath.Dir(filePath), 0o755)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0o644)
		if err == nil {
			_, _ = f.Write(content)
		}
	} else {
		_ = os.WriteFile(filePath, content, 0o644)
	}
}

func StorageCreateMultiFile(_, _ string, _ *http.Request, Bucket string, fList []struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
},
) {
	for _, s := range fList {
		content, _ := json.Marshal(s.Data)
		fileName := s.Name
		thisTmp := "/var/www/autoketing-storage/" + Bucket
		filePath := thisTmp + "/" + fileName
		_ = os.MkdirAll(filepath.Dir(filePath), 0o755)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0o644)
			if err == nil {
				_, _ = f.Write(content)
			}
		} else {
			_ = os.WriteFile(filePath, content, 0o644)
		}
	}
}

func StorageDeleteFile(_, _ string, _ *http.Request, _ string, _ string) {
	return
}

func StorageReadFile(
	_, _ string,
	_ *http.Request,
	_ string,
	_ string,
) ([]byte, error) {
	return nil, nil
}

func StorageCheckFile(_, _ string, _ *http.Request, Bucket string, fileName string) bool {
	var checked bool
	thisTmp := "/var/www/autoketing-storage/" + Bucket
	filePath := thisTmp + "/" + fileName
	_ = os.MkdirAll(filepath.Dir(filePath), 0o755)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		checked = true
	}
	return checked
}

func StorageDeleteMultiFile(
	_, _ string,
	_ *http.Request,
	_ string,
	_ []string,
) {
	return
}
