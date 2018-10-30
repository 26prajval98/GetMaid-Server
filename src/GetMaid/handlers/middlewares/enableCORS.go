package middlewares

import "net/http"

func EnableCors(res http.ResponseWriter, req *http.Request) bool {
	_ = req
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	return true
}
