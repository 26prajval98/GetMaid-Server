package dblogin

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"strconv"
)

func post(req *http.Request, res http.ResponseWriter) {

	var databaseEmail string
	var databasePassword string

	// 1.Username and password
	var err error
	defer methods.ErrorHandler(res, &err)

	email := req.FormValue("Email")
	password := req.FormValue("Password")

	db := database.GetDb()

	err = db.QueryRow("SELECT Password FROM hirer WHERE Name=?", email).Scan(&databasePassword)

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))

	if req.Form.Get("Password") != databasePassword {
		panic("Incorrect password")
	}

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	if err == nil {
		io.WriteString(res, "Hello "+databaseEmail)
		success, err := json.Marshal(types.Success{Success: true, Msg: strconv.Itoa(NOERROR)})
		methods.CheckErr(err)
		methods.SendJSONResponse(res, success, 200)
	}

}
