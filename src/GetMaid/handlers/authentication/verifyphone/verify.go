package verifyphone

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func post(res http.ResponseWriter, req *http.Request) {
	var err error

	methods.ErrorHandler(res, &err)

	req.ParseForm()

	ph := req.Header.Get("Phone")
	otp := req.FormValue("otp")
	isMaid, _ := strconv.Atoi(req.Header.Get("Maid"))
	secpexiry := req.FormValue("Secret")
	s := secpexiry[5:]
	exp, err := strconv.Atoi(secpexiry[:5])

	if err != nil {
		panic(7)
	}

	current := time.Now().Unix()

	if (current/100000)*100000+int64(exp) < current {
		panic("OTP expired")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(s), []byte(otp)); err != nil {
		panic("Incorrect otp")
		return
	}

	db := database.GetDb()

	if isMaid == 1 {
		_, err = db.Exec(`Update maid
						Set Active = 1
						Where Phone = ?
						`, ph)

		if err != nil {
			panic(7)
		}
	} else {
		_, err = db.Exec(`Update hirer
						Set Active = 1
						Where Phone = ?
						`, ph)

		if err != nil {
			panic(7)
		}
	}
	success, err := json.Marshal(types.Success{Success: true})
	methods.CheckErr(err)
	methods.SendJSONResponse(res, success, 200)
}

func Handler(res http.ResponseWriter, req *http.Request) error {
	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("POST", "/verify", req):
		post(res, req)
	default:
		panic(404)
	}

	return e
}
