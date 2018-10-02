package handlers

import (
	"GetMaid/handlers/methods"
	"io"
	"net/http"
)

func SignupGetHandler(res http.ResponseWriter) {
	io.WriteString(res, "Signup Route")
	wg.Done()
}

func SignupPostHandler(req *http.Request, res http.ResponseWriter) {

	var e error

	defer methods.ErrorHandler(res, &e, &wg)

	req.ParseForm()

	if len(req.Form["username"][0]) == 0 || len(req.Form["password"]) == 0 {
		panic("USERNAME OR PASSWORD MISSING")
	}

	wg.Done()
}

func SignupHandler(res http.ResponseWriter, req *http.Request) error {
	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("GET", "/signup", req):
		wg.Add(1)
		go SignupGetHandler(res)
	case methods.CheckCase("POST", "/signup", req):
		wg.Add(1)
		go SignupPostHandler(req, res)
	default:
		panic(404)
	}

	wg.Wait()

	return e
}
