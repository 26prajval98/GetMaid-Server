package dblogin



import (
	"GetMaid/handlers/methods"
	"net/http"
)

const NOERROR  = iota
func Handler(res http.ResponseWriter, req *http.Request) error {
	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("GET", "/login", req):
		get(res)

	case methods.CheckCase("POST", "/login", req):
		post(req, res)
	default:
		panic(404)
	}

	return e
}
