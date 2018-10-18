package main

import (
	"GetMaid/database"
	"GetMaid/handlers/authentication/jwt"
	"GetMaid/handlers/authentication/local"
	"GetMaid/handlers/authentication/verifyphone"
	"GetMaid/handlers/index"
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
	server.HandlePath("/login", local.Handler)
	server.HandlePath("/verify", verifyphone.Handler, jwt.VerifyJWT)

	fmt.Println("Server Started")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
