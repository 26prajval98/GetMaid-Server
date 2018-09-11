package handlers

import (
	"io"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func getHandler(resp http.ResponseWriter) {
	io.WriteString(resp, "Welcome to GetMaid")
	wg.Done()
}

func IndexHandler(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "Get":
		wg.Add(1)
		go getHandler(resp)
	default:
		wg.Add(1)
		go getHandler(resp)
	}
	wg.Wait()
}
