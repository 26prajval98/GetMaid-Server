package methods

import (
	"net/http"
)

func SendJSONResponse(res http.ResponseWriter, s []byte, code int) {
	res.WriteHeader(code)
	res.Header().Set("Content-Type", "application/json")
	res.Write(s)
}
