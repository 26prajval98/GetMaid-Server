package signup

import (
	"io"
	"net/http"
)

func get(res http.ResponseWriter) {
	io.WriteString(res, "Signup Route")
}
