package go_gcloud_engine

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func AddNewTaskQueue(
	Env string,
	request *http.Request,
	path string,
	param url.Values,
	queueName string,
) (task interface{}, err error) {
	u, _ := url.ParseRequestURI(QueueService)
	u.Path = fmt.Sprintf("/add-taskqueue/%s", queueName)
	urlStr := u.String()

	var client = &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(param.Encode())) // URL-encoded payload
	r.Header.Add("Authorization", fmt.Sprintf("auth_token=\"%s\"", Env))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("taskqueue", path)
	r.Header.Add("request", request.Host)

	_, err = client.Do(r)
	return
}
