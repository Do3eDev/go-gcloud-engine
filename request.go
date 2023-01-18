package go_gcloud_engine

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func WriteLogError(Env string, request *http.Request, format string, err interface{}) {
	if Env == "local" {
		fmt.Println(format, err)
	} else {
		//ctx := appengine.NewContext(request)
		//log.Errorf(ctx, format, err)
	}
}

func RequestCustomer(
	Env string,
	method string,
	url string,
	body []byte,
	header map[string]string,
	request *http.Request,
) (status int, responseBody []byte, err error, Header map[string][]string) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		status = 1
		return
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}
	var resp *http.Response

	client := http.DefaultClient
	resp, err = client.Do(req)

	if err != nil {
		status = 2
		return
	}

	defer resp.Body.Close()
	status = resp.StatusCode
	Header = resp.Header
	responseBody, err = ioutil.ReadAll(resp.Body)
	return
}
