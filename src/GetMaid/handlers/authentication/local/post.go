package local

import (
	"GetMaid/database"
	"GetMaid/handlers/authentication/jwt"
	"GetMaid/handlers/authentication/verifyphone"
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const (
	INVALIDLOGIN = 6
)

func setCookie(msg []byte, res *http.ResponseWriter) {
	http.SetCookie(*res, &http.Cookie{Name: "token", Value: string(msg), Expires: time.Now().Add(24 * time.Hour)})
}

func post(req *http.Request, res http.ResponseWriter) {

	var databasePhone string
	var databasePassword string
	var databaseName string
	var databaseActive int
	var id int64

	// 1.Username and password
	var err error
	var check bool
	defer methods.ErrorHandler(res, &err)

	req.ParseForm()

	emailOrPhone := req.FormValue("EmailOrPhone")
	password := req.FormValue("Password")
	isMaid, _ := strconv.Atoi(req.FormValue("IsMaid"))

	db := database.GetDb()

	if isMaid == 0 {
		if check, _ = regexp.MatchString(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, emailOrPhone); !check {
			err := db.QueryRow("SELECT Hirer_id, Name, Phone, Password, Active FROM hirer WHERE Phone=?", emailOrPhone).Scan(&id, &databaseName, &databasePhone, &databasePassword, &databaseActive)
			if err != nil {
				panic(INVALIDLOGIN)
				return
			}
		} else {
			err = db.QueryRow("SELECT Hirer_id, Name, Phone, Password, Active FROM hirer WHERE Email=?", emailOrPhone).Scan(&id, &databaseName, &databasePhone, &databasePassword, &databaseActive)
			if err != nil {
				panic(INVALIDLOGIN)
				return
			}
		}
	} else {
		if check, _ = regexp.MatchString(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, emailOrPhone); !check {
			err = db.QueryRow("SELECT Maid_id, Name, Phone, Password, Active FROM maid WHERE Phone=?", emailOrPhone).Scan(&id, &databaseName, &databasePhone, &databasePassword, &databaseActive)
			if err != nil {
				panic(INVALIDLOGIN)
				return
			}
		} else {
			err = db.QueryRow("SELECT Maid_id, Name, Phone, Password, Active FROM maid WHERE Email=?", emailOrPhone).Scan(&id, &databaseName, &databasePhone, &databasePassword, &databaseActive)
			if err != nil {
				panic(INVALIDLOGIN)
				return
			}
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))

	if err != nil {
		panic(INVALIDLOGIN)
	}

	msg, err := jwt.GenerateJWT(id, databasePhone, methods.IntToBool(isMaid))

	if err != nil {
		panic(msg)
	}

	fmt.Println(databaseActive)

	setCookie(msg, &res)

	if databaseActive == 1 {
		success, err := json.Marshal(types.Success{Success: true, Msg: string(msg)})
		res.Header().Add("maid", strconv.Itoa(isMaid))
		methods.CheckErr(err)
		methods.SendJSONResponse(res, success, 200)
	} else {
		req.Header.Set("Msg", string(msg))
		req.Header.Set("Phone", databasePhone)
		req.Header.Set("Maid", strconv.Itoa(isMaid))
		verifyphone.Send(res, req)
	}

}
