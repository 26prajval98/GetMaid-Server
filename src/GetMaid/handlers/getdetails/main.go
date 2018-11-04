package getdetails

import (
	"GetMaid/handlers/methods"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) error {
	var e error
	defer methods.ErrorHandler(res, &e)
	_ = req
	switch {
	case methods.CheckCase("GET", "/details", req):
		get(req, res)
	case methods.CheckCase("POST", "/details", req):
		post(req)
	}
	return nil
}
