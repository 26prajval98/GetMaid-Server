package main

import (
	"GetMaid/database"
	"GetMaid/handlers/authentication/jwt"
	"GetMaid/handlers/authentication/local"
	"GetMaid/handlers/authentication/verifyphone"
	"GetMaid/handlers/getdetails"
	"GetMaid/handlers/index"
	"GetMaid/handlers/ismaid"
	"GetMaid/handlers/maidservices"
	"GetMaid/handlers/middlewares"
	"GetMaid/handlers/signup"
	"GetMaid/server"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {

	defer database.CloseDb()

	mux := http.NewServeMux()

	server.HandlePath("/", mux, index.Handler, jwt.VerifyJWT)
	server.HandlePath("/signup", mux, signup.Handler)
	server.HandlePath("/login", mux, local.Handler, middlewares.EnableCors)
	server.HandlePath("/verify", mux, verifyphone.Handler, jwt.VerifyJWT)

	//Common
	server.HandlePath("/details", mux, getdetails.Handler, jwt.VerifyJWT)
	server.HandlePath("/ismaid", mux, ismaid.Handler, jwt.VerifyJWT)

	// Maid Paths
	server.HandlePath("/maidservices", mux, maidservices.Handler, jwt.VerifyJWT, middlewares.IsMaid)

	fmt.Println("Server Started")

	handler := cors.AllowAll().Handler(mux)

	log.Fatal(http.ListenAndServe(":3000", handler))
}
