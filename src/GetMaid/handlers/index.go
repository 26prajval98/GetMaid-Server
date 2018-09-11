package handlers

import (
	"io"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func getHandler(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello Go")
	wg.Done()
}

func IndexHandler(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "Get":
		wg.Add(1)
		go getHandler(resp, req)
	default:
		wg.Add(1)
		go getHandler(resp, req)
	}
	wg.Wait()
}
