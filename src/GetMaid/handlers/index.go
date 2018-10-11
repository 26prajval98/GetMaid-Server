package handlers

import (
	"GetMaid/handlers/methods"
	"io"
	"net/http"

)



func indexGetHandler(res http.ResponseWriter) {
	io.WriteString(res, "Welcome to GetMaid")

}

func IndexHandler(res http.ResponseWriter, req *http.Request) error {

	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("GET", "/", req):
		 indexGetHandler(res)
	default:
		panic(404)
	}

	return nil
}
