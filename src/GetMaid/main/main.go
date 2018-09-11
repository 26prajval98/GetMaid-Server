package main

import (
	"GetMaid/handlers"
	"GetMaid/server"
	"log"
	"net/http"
)

func main() {
	server.HandlePath("/", handlers.IndexHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
