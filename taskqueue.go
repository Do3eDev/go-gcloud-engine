package go_gcloud_engine

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/variar/buckets"
)

func AddNewTaskQueue(
	_ string,
	request *http.Request,
	path string,
	param url.Values,
	queueName string,
) (task interface{}, err error) {
	if strings.TrimSpace(queueName) == "" {
		queueName = "default"
	}
	folder := fmt.Sprintf("queue/%s", time.Now().Format(time.DateOnly))
	_ = os.MkdirAll(folder, 0o755)
	bx, _ := buckets.Open(fmt.Sprintf("%s/%s.queue", folder, queueName))
	defer bx.Close()
	todos, _ := bx.New([]byte("todos"))
	sb1, err := io.ReadAll(strings.NewReader(param.Encode()))
	if err == nil {
		_ = todos.Put(
			[]byte(fmt.Sprintf("%d|||https://%s%s", time.Now().UnixNano(), request.Host, path)),
			sb1,
		)
	}
	return
}
