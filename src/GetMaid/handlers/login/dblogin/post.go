package dblogin

import (
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"strconv"
)


func post(req *http.Request,res http.ResponseWriter){

	// 1.Username and password
	var err error
	defer methods.ErrorHandler(res, &err)


	username:=req.FormValue("username")
	password:=req.FormValue("password")

	var databaseUsername string
	var databasePassword string

	db, e := sql.Open("mysql", "root:namah1998@tcp(127.0.0.1:3306)/getmaid")
	if e != nil {
		log.Fatal("Database Not Connected")
	}

	err = db.QueryRow("SELECT Name,Password FROM hirer WHERE Name=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))

	if req.Form.Get("Password") != databasePassword {
		panic(PASSWORD)
	}

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}


	if err==nil{
		io.WriteString(res,"Hello "+databaseUsername)
		success, err := json.Marshal(types.Success{Success: true, Msg: strconv.Itoa(NOERROR)})
		methods.CheckErr(err)
		methods.SendJSONResponse(res, success, 200)
	}






}
