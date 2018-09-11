package server

import (
	"GetMaid/logger"
	"io"
	"net/http"
	"time"
)

func errorHandler(err error, res http.ResponseWriter, info ...interface{}) bool {

	if err != nil {
		logger.WarnLog(err, info)
		io.WriteString(res, "404 Page Not Found")
		return true
	}
	return false
}

func HandlePath(path string, handler func(http.ResponseWriter, *http.Request) error, middlewares ...func()) {

	http.HandleFunc(path, func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		for _, f := range middlewares {
			f()
		}

		err := handler(res, req)

		elapsed := time.Since(startTime).String()

		if !errorHandler(err, res, req.Method, path) {
			logger.InfoLog(req.Method, path, elapsed)
		}
	})
}
