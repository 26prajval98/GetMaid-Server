package handlers

import (
	"GetMaid/handlers/methods"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

var db *sql.DB


func LoginGetHandler(res http.ResponseWriter){
	io.WriteString(res, "Login Route")
	//http.ServerFile(res,req,"login.html")
}

func LoginPostHandler(req *http.Request,res http.ResponseWriter){

	// 1.Username and password

	username:=req.FormValue("username")
	password:=req.FormValue("password")

	var databaseUsername string
	var databasePassword string

	err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	io.WriteString(res,"Hello "+databaseUsername)

	//Facebook oauth




}

func LoginHandler(res http.ResponseWriter, req *http.Request) error {
	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("GET", "/login", req):
		LoginGetHandler(res)

	case methods.CheckCase("POST", "/login", req):

		LoginPostHandler(req, res)
	default:
		panic(404)
	}

	return e
}
//func Login(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("method:", r.Method) //get request method
//	if r.Method == "GET" {
//		t, _ := template.ParseFiles("login.gtpl")
//		t.Execute(w, nil)
//	} else {
//		r.ParseForm()
//		// logic part of log in
//		fmt.Println("username:", r.Form["username"])
//		fmt.Println("password:", r.Form["password"])
//	}
//}
