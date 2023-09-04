package tools

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// WriteHttpJsonResponse 向HTTP中写入Json结果
func WriteHttpJsonResponse(rwPtr *http.ResponseWriter, data any) {
	rw := *rwPtr
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		logrus.Errorf("写入ResponseBody错误，errMsg=%s", err.Error())
		io.WriteString(rw, err.Error())
	}
}

// CheckHttpMethod 检查HttpMethod是否合法
func CheckHttpMethod(req *http.Request, methods ...string) bool {
	res := false
	for _, m := range methods {
		if res {
			break
		}
		res = req.Method == m
	}
	return res
}
