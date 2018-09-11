package handlers

import (
	"github.com/pkg/errors"
	"io"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func getHandler(res http.ResponseWriter) {
	io.WriteString(res, "Welcome to GetMaid")
	wg.Done()
}

func IndexHandler(res http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		wg.Add(1)
		go getHandler(res)
	default:
		return errors.New("Page not found")
	}
	wg.Wait()

	return nil
}
