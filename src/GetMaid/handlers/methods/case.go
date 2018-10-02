package methods

import "net/http"

func CheckCase(method string, path string, RQ *http.Request) (y bool) {
	y = method == RQ.Method && path == RQ.URL.Path
	return
}
