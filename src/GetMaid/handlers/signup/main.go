package signup

import (
	"GetMaid/handlers/methods"
	"net/http"
)

const (
	NOERROR  = iota
	PASSWORD = iota
	EMAIL    = iota
	PHONE    = iota
	NAME     = iota
	ADDRESS
)

type Maid struct {
	Email    string
	Password string
	Phone    string
	Name     string
	Address  Address
}

type Hirer struct {
	Email    string
	Password string
	Phone    string
	Name     string
	HouseNo  string
	Address  Address
}

type Address struct {
	Locality string
	City     string
	PinCode  string
}

func Handler(res http.ResponseWriter, req *http.Request) error {
	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("GET", "/signup", req):
		get(res)

	case methods.CheckCase("POST", "/signup", req):
		post(req, res)

	default:
		panic(404)
	}

	return e
}
