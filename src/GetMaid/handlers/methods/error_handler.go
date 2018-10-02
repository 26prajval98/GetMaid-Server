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
			}
		}

		s, e := json.Marshal(types.Success{Success: false, Msg: msg})

		if e != nil {
			log.Fatal(e)
			res.WriteHeader(500)
		}

		//fmt.Println(code)

		*err = errors.New(msg)

		res.WriteHeader(code)
		res.Header().Set("Content-Type", "application/json")
		res.Write(s)

		if len(wg) > 0 {
			wg[0].Done()
		}
	}
}
