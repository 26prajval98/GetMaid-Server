package maidservices

import (
	"GetMaid/handlers/methods"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) error {

	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("GET", "/maidservices", req):
		get(res, req.Header.Get("Maid_id"))
	case methods.CheckCase("POST", "/maidservices", req):
		req.ParseForm()
		post(res, req.Header.Get("Maid_id"), req.Form.Get("Service_name"))
	case methods.CheckCase("DELETE", "/maidservices", req):
		q := req.URL.Query()
		del(res, q.Get("i"))
	default:
		panic(404)
	}

	return nil
}
