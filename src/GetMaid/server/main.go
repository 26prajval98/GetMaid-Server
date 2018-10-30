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

func HandlePath(path string, mux *http.ServeMux, handler func(http.ResponseWriter, *http.Request) error, middlewares ...func(http.ResponseWriter, *http.Request) bool) {
	mux.HandleFunc(path, func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		var err error

		func() {
			for _, f := range middlewares {
				next := f(res, req)

				if !next {
					return
				}
			}

			err = handler(res, req)
		}()

		elapsed := time.Since(startTime).String()

		if !errorHandler(err, res, req.Method, path) {
			logger.InfoLog(req.Method, path, elapsed)
		}
	})
}
