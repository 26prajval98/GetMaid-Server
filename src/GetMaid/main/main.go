package main

import (
	"GetMaid/database"
	"GetMaid/handlers/index"
	"GetMaid/handlers/login/dblogin"
	"GetMaid/handlers/login/googlelogin"
	"GetMaid/handlers/signup"
	"GetMaid/server"
	"fmt"
	"log"
	"net/http"
)

func main() {

	defer database.CloseDb()
	server.HandlePath("/", index.Handler)
	server.HandlePath("/signup", signup.Handler)
	server.HandlePath("/login",dblogin.Handler)
	server.HandlePath("/login/google_login",googlelogin.Handler)
	server.HandlePath("/login/google_login/loginpage",googlelogin.LoginHandler)
	server.HandlePath("login/google_login/callback",googlelogin.CallbackHandler)

	fmt.Println("Server Started")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
