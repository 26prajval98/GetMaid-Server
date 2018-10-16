package index

import (
	"GetMaid/handlers/methods"
	"io"
	"net/http"
)

func get(res http.ResponseWriter, msg string) {
	io.WriteString(res, "Welcome to GetMaid "+msg)
}

func Handler(res http.ResponseWriter, req *http.Request) error {

	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("GET", "/", req):
		get(res, req.Header.Get("Maid_id")+" "+req.Header.Get("Phone"))
	default:
		panic(404)
	}

	return nil
}
