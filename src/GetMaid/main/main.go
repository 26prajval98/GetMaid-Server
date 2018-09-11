package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, "Hello Go")
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
