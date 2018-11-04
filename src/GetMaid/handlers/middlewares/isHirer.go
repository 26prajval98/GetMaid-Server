package middlewares

import (
	"net/http"
	"strconv"
)

func IsHirer(res http.ResponseWriter, req *http.Request) (next bool) {
	i, err := strconv.Atoi(req.Header.Get("Maid"))
	if err != nil || i == 1 {
		errorHandler(res, &next, "Not Hirer")
	}

	next = true
	return
}
