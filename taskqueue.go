package go_gcloud_engine

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/taskqueue"
	"net/http"
	"net/url"
)

func AddNewTaskQueue(Env string, request *http.Request, path string, param url.Values, queueName string) (task *taskqueue.Task, err error) {
	if Env == "local" {
		return
	}
	return taskqueue.Add(
		appengine.NewContext(request),
		taskqueue.NewPOSTTask(path, param),
		queueName)
}
