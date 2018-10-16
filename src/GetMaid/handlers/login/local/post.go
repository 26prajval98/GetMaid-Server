package local

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"strconv"
)

func post(req *http.Request, res http.ResponseWriter) {

	var databaseEmail string
	var databasePassword string
	var databaseName string

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
			err := db.QueryRow("SELECT Name, Email, Password FROM hirer WHERE Phone=?", emailOrPhone).Scan(&databaseName, &databaseEmail, &databasePassword)
			if err != nil {
				panic("Incorrect email / phone or password")
				return
			}
		} else {
			err = db.QueryRow("SELECT Name, Email, Password FROM hirer WHERE Email=?", emailOrPhone).Scan(&databaseName, &databaseEmail, &databasePassword)
			if err != nil {
				panic("Incorrect email / phone or password")
				return
			}
		}
	} else {
		if check, _ = regexp.MatchString(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, emailOrPhone); !check {
			err = db.QueryRow("SELECT Name, Email, Password FROM maid WHERE Phone=?", emailOrPhone).Scan(&databaseName, &databaseEmail, &databasePassword)
			if err != nil {
				panic("Incorrect email / phone or password")
				return
			}
		} else {
			err = db.QueryRow("SELECT Name, Email, Password FROM maid WHERE Email=?", emailOrPhone).Scan(&databaseName, &databaseEmail, &databasePassword)
			if err != nil {
				panic("Incorrect email / phone or password")
				return
			}
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))

	if err != nil {
		panic(err.Error())
	}

	success, err := json.Marshal(types.Success{Success: true, Msg: databaseName})
	methods.CheckErr(err)
	methods.SendJSONResponse(res, success, 200)
}
