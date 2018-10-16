package local

import (
	"io"
	"net/http"
)

func get(res http.ResponseWriter) {
	io.WriteString(res, "Login Route")
}
