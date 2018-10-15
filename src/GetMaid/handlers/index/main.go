package index

import (
	"GetMaid/handlers/methods"
	"io"
	"net/http"
)

func get(res http.ResponseWriter) {
	io.WriteString(res, "Welcome to GetMaid")
}

func Handler(res http.ResponseWriter, req *http.Request) error {

	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("GET", "/", req):
		get(res)
	default:
		panic(404)
	}

	return nil
}
