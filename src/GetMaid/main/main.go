package main

import (
	"GetMaid/handlers/index"
	"GetMaid/handlers/signup"
	"GetMaid/server"
	"fmt"
	"log"
	"net/http"
)

func main() {

	server.HandlePath("/", index.Handler)
	server.HandlePath("/signup", signup.Handler)

	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
