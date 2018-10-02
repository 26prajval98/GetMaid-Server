package main

import (
	"GetMaid/handlers"
	"GetMaid/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server.HandlePath("/", handlers.IndexHandler)

	server.HandlePath("/signup", handlers.SignupHandler)

	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
