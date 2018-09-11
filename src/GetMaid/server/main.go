package server

import (
	"GetMaid/logger"
	"net/http"
	"time"
)

func HandlePath(msg string, handler func(http.ResponseWriter, *http.Request), middlewares ...func()) {

	http.HandleFunc(msg, func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		for _, f := range middlewares {
			f()
		}

		handler(res, req)

		elapsed := time.Since(startTime).String()

		logger.Infolog(elapsed)
	})
}
