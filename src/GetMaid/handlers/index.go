package handlers

import (
	"GetMaid/handlers/methods"
	"io"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func indexGetHandler(res http.ResponseWriter) {
	io.WriteString(res, "Welcome to GetMaid")
	wg.Done()
}

func IndexHandler(res http.ResponseWriter, req *http.Request) error {

	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("GET", "/", req):
		wg.Add(1)
		go indexGetHandler(res)
	default:
		panic(404)
	}
	wg.Wait()

	return nil
}
