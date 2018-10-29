package main

import (
	"GetMaid/database"
	"GetMaid/handlers/authentication/jwt"
	"GetMaid/handlers/authentication/local"
	"GetMaid/handlers/authentication/verifyphone"
	"GetMaid/handlers/index"
	"GetMaid/handlers/maidservices"
	"GetMaid/handlers/middlewares"
	"GetMaid/handlers/signup"
	"GetMaid/server"
	"fmt"
	"log"
	"net/http"
)

func main() {

	defer database.CloseDb()
	server.HandlePath("/", index.Handler, jwt.VerifyJWT)
	server.HandlePath("/signup", signup.Handler)
	server.HandlePath("/login", local.Handler)
	server.HandlePath("/verify", verifyphone.Handler, jwt.VerifyJWT)
	server.HandlePath("/maidservices", maidservices.Handler, jwt.VerifyJWT, middlewares.IsMaid)

	fmt.Println("Server Started")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
