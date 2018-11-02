package maidsearch

import (
	"GetMaid/handlers/methods"
	"net/http"
)

func Handler(res http.ResponseWriter,req *http.Request)error{
	var e error
	defer methods.ErrorHandler(res, &e)

	switch {

	case methods.CheckCase("GET", "/maidsearch", req):
		search(req, res)

	default:
		panic(404)
	}

	return e

}
