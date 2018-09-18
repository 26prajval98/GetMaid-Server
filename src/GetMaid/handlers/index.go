package handlers

import (
	"github.com/pkg/errors"
	"io"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

var path = "/"

var RQ *http.Request

func getHandler(res http.ResponseWriter) {
	io.WriteString(res, "Welcome to GetMaid")
	wg.Done()
}

func checkCase(method string) (y bool) {
	y = method == RQ.Method && path == RQ.URL.Path
	return
}

func IndexHandler(res http.ResponseWriter, req *http.Request) error {
	RQ = req
	switch {
	case checkCase("GET"):
		wg.Add(1)
		go getHandler(res)
	default:
		return errors.New("Page not found")
	}
	wg.Wait()

	return nil
}
