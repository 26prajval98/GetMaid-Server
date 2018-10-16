package dblogin

import (
	"io"
	"net/http"
)

func get(res http.ResponseWriter){
	io.WriteString(res, "Login Route")
	//http.ServerFile(res,req,"login.html")
}

