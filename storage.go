package go_gcloud_engine

import (
	"bytes"
	"cloud.google.com/go/storage"
	"compress/gzip"
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

var ctx = context.Background()

func StorageCreateFile(Env string, request *http.Request, Bucket, fileName string, content []byte) {
	if Env == "local" {
		return
	}

	var client, err = storage.NewClient(ctx)
	if err != nil {
		return
	} else {
		defer client.Close()
	}
	var buf = &bytes.Buffer{}
	var ws = client.Bucket(Bucket).Object(fileName).NewWriter(ctx)
	ws.ContentType = "application/json"
	ws.ContentEncoding = "gzip"
	var zw = gzip.NewWriter(buf)
	_, _ = zw.Write(content)
	_ = zw.Close()
	_, _ = io.Copy(ws, buf)
	_ = ws.Close()
}

func StorageCreateFileSVG(Env string, request *http.Request, Bucket, fileName string, content []byte) {
	if Env == "local" {
		return
	}

	var client, err = storage.NewClient(ctx)
	if err != nil {
		return
	} else {
		defer client.Close()
	}
	var buf = &bytes.Buffer{}
	var ws = client.Bucket(Bucket).Object(fileName).NewWriter(ctx)
	ws.ContentType = "image/svg+xml"
	ws.ContentEncoding = "gzip"
	var zw = gzip.NewWriter(buf)
	_, _ = zw.Write(content)
	_ = zw.Close()
	_, _ = io.Copy(ws, buf)
	_ = ws.Close()
}

func StorageCreateMultiFile(Env string, request *http.Request, Bucket string, fList []struct {
	Name string
	Data []byte
}) {
	if Env == "local" {
		return
	}

	var client, err = storage.NewClient(ctx)
	if err != nil {
		return
	} else {
		defer client.Close()
	}
	for _, v := range fList {
		var buf = &bytes.Buffer{}
		var ws = client.Bucket(Bucket).Object(v.Name).NewWriter(ctx)
		ws.ContentType = "application/json"
		ws.ContentEncoding = "gzip"
		var zw = gzip.NewWriter(buf)
		_, _ = zw.Write(v.Data)
		_ = zw.Close()
		_, _ = io.Copy(ws, buf)
		_ = ws.Close()
	}
}

func StorageDeleteFile(Env string, request *http.Request, Bucket string, fileName string) {
	if Env == "local" {
		return
	}

	var client, err = storage.NewClient(ctx)
	if err != nil {
		return
	} else {
		defer client.Close()
	}
	_ = client.Bucket(Bucket).Object(fileName).Delete(ctx)
	return
}

func StorageReadFile(Env string, request *http.Request, Bucket string, fileName string) ([]byte, error) {

	var client, err = storage.NewClient(ctx)
	if err != nil {
		return nil, err
	} else {
		defer client.Close()
	}
	rc, err := client.Bucket(Bucket).Object(fileName).NewReader(ctx)
	if err != nil {
		return nil, err
	} else {
		defer rc.Close()
	}
	return ioutil.ReadAll(rc)
}

func StorageCheckFile(Env string, request *http.Request, Bucket string, fileName string) bool {
	if Env == "local" {
		return false
	}

	var client, err = storage.NewClient(ctx)
	if err != nil {
		return false
	} else {
		defer client.Close()
	}
	_, err = client.Bucket(Bucket).Object(fileName).NewReader(ctx)
	if err != nil {
		return false
	}
	return true
}

func StorageDeleteMultiFile(Env string, request *http.Request, Bucket string, array1 []string) {
	if Env == "local" {
		return
	}

	var client, err = storage.NewClient(ctx)
	if err != nil {
		return
	} else {
		defer client.Close()
	}
	for _, v := range array1 {
		_ = client.Bucket(Bucket).Object(v).Delete(ctx)
	}
	return
}
