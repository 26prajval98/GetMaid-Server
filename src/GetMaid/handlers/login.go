package handlers

import (
	"GetMaid/handlers/methods"
	"database/sql"
	"github.com/qor/auth"
	//"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/providers/facebook"
	"github.com/qor/auth/providers/google"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

//var db *sql.DB
var (

	// Initialize SQL DB
	sqlDB, _ = sql.Open("mysql","getmaid.db")

	// Initialize Auth with configuration
	Auth = auth.New(&auth.Config{
		DB: sqlDB,
	})
)

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

	err := sqlDB.QueryRow("SELECT username, password FROM hirer WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}
	if err==nil{
		io.WriteString(res,"Hello "+databaseUsername)
	}

	//// Migrate AuthIdentity model, AuthIdentity will be used to save auth info, like username/password, oauth token, you could change that.
	//sqlDB.AutoMigrate(&auth_identity.AuthIdentity{})


	// 2.Allow use Google
	Auth.RegisterProvider(google.New(&google.Config{
		ClientID:     "google client id",
		ClientSecret: "google client secret",
	}))

	// 2.Allow use Facebook
	Auth.RegisterProvider(facebook.New(&facebook.Config{
		ClientID:     "facebook client id",
		ClientSecret: "facebook client secret",
	}))



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

