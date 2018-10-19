package methods

import (
	"GetMaid/handlers/types"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
)

func ErrorHandler(res http.ResponseWriter, err *error, wg ...*sync.WaitGroup) {

	*err = nil

	if r := recover(); r != nil {
		var (
			msg  string
			code int
		)

		switch r {

		case 404:
			code = 404
			msg = "page not found"

		case 500:
			code = 500
			msg = "internal server error"

		default:
			code = 400
			switch x := r.(type) {

			case string:
				msg = x
			case int:
				switch x {
				case 1:
					msg = "Passwords Do Not Match"
				case 2:
					msg = "Invalid Email"
				case 3:
					msg = "Invalid Phone Number"
				case 4:
					msg = "Invalid Name"
				case 5:
					msg = "Invalid Address"
				case 6:
					msg = "Incorrect Email Or Passowrd"
				case 7:
					msg = "Something went wrong"
				}
			}
		}

		s, e := json.Marshal(types.Success{Success: false, Msg: msg})

		if e != nil {
			log.Fatal(e)
			res.WriteHeader(500)
		}

		*err = errors.New(msg)

		SendJSONResponse(res, s, code)

		if len(wg) > 0 {
			wg[0].Done()
		}

	}
}
