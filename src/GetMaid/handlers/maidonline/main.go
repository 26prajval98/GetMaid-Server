package maidonline

import (
	"GetMaid/handlers/methods"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) error {
	var e error
	defer methods.ErrorHandler(res, &e)
	_ = req
	switch {
	case methods.CheckCase("GET", "/maidonline", req):
		get(req)
	case methods.CheckCase("POST", "/maidonline", req):
		getAdd(req)
	case methods.CheckCase("DELETE", "/maidonline", req):
		del(req)
	}
	return nil
}
