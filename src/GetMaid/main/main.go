package main

import (
	"GetMaid/database"
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

	fmt.Println("Server Started")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
