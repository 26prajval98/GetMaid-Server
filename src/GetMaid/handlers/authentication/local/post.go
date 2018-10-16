package local

import (
	"GetMaid/database"
	"GetMaid/handlers/authentication/jwt"
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"strconv"
)

const (
	INVALIDLOGIN = 6
)

func post(req *http.Request, res http.ResponseWriter) {

	var databasePhone string
	var databasePassword string
	var databaseName string
	var id int64

	// 1.Username and password
	var err error
	var check bool
	defer methods.ErrorHandler(res, &err)

	emailOrPhone := req.FormValue("EmailOrPhone")
	password := req.FormValue("Password")
	isMaid, _ := strconv.Atoi(req.FormValue("IsMaid"))

	db := database.GetDb()

	if isMaid == 0 {
		if check, _ = regexp.MatchString(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, emailOrPhone); !check {
			err := db.QueryRow("SELECT Hirer_id, Name, Phone, Password FROM hirer WHERE Phone=?", emailOrPhone).Scan(&id, &databaseName, &databasePhone, &databasePassword)
			if err != nil {
				panic(INVALIDLOGIN)
				return
			}
		} else {
			err = db.QueryRow("SELECT Hirer_id, Name, Phone, Password FROM hirer WHERE Email=?", emailOrPhone).Scan(&id, &databaseName, &databasePhone, &databasePassword)
			if err != nil {
				panic(INVALIDLOGIN)
				return
			}
		}
	} else {
		if check, _ = regexp.MatchString(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, emailOrPhone); !check {
			err = db.QueryRow("SELECT Maid_id, Name, Phone, Password FROM maid WHERE Phone=?", emailOrPhone).Scan(&id, &databaseName, &databasePhone, &databasePassword)
			if err != nil {
				panic(INVALIDLOGIN)
				return
			}
		} else {
			err = db.QueryRow("SELECT Maid_id, Name, Phone, Password FROM maid WHERE Email=?", emailOrPhone).Scan(&id, &databaseName, &databasePhone, &databasePassword)
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

	success, err := json.Marshal(types.Success{Success: true, Msg: string(msg)})
	methods.CheckErr(err)
	methods.SendJSONResponse(res, success, 200)
}
