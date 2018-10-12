package signup

import (
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func post(req *http.Request, res http.ResponseWriter) {

	var e error

	defer methods.ErrorHandler(res, &e)

	req.ParseForm()

	if len(req.Form.Get("Password")) < 10 {
		panic("Minimum length of password is 10 letters")
	}

	if req.Form.Get("Password") != req.Form.Get("Repassword") {
		panic(PASSWORD)
	}

	hpw, e := bcrypt.GenerateFromPassword([]byte(req.Form.Get("Password")), 6)

	if e != nil {
		panic(500)
	}

	if t, _ := strconv.Atoi(req.Form.Get("IsMaid")); t == int(0) {
		user := Hirer{
			Email:    req.Form.Get("Email"),
			Password: string(hpw),
			Phone:    req.Form.Get("Phone"),
			Name:     req.Form.Get("Name"),
			HouseNo:  req.Form.Get("HouseNo"),
			Address: Address{
				Locality: req.Form.Get("Locality"),
				PinCode:  req.Form.Get("PinCode"),
				City:     req.Form.Get("City"),
			},
		}
		validateSignup(user)
	} else {
		user := Maid{
			Email:    req.Form.Get("Email"),
			Password: string(hpw),
			Phone:    req.Form.Get("Phone"),
			Name:     req.Form.Get("Name"),
			Address: Address{
				Locality: req.Form.Get("Locality"),
				PinCode:  req.Form.Get("PinCode"),
				City:     req.Form.Get("City"),
			},
		}
		validateSignup(user)
	}

	success, err := json.Marshal(types.Success{Success: true, Msg: strconv.Itoa(NOERROR)})
	methods.CheckErr(err)
	methods.SendJSONResponse(res, success, 200)
}
