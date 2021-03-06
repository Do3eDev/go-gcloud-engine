package go_gcloud_engine

import (
	"bytes"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
	"net/http"
)

func SendgridRequestSendMail(Env string, method string, url string, body []byte, header map[string]string, e *http.Request) (statusCode int, messageId string, rspBody []byte, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	var resp *http.Response
	if Env != "local" {
		resp, err = urlfetch.Client(appengine.NewContext(e)).Do(req)
	} else {
		resp, err = http.DefaultClient.Do(req)
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	messageId = resp.Header.Get("X-Message-Id")
	rspBody, err = ioutil.ReadAll(resp.Body)
	return
}
