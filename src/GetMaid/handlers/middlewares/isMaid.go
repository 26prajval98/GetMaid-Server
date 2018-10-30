package middlewares

import (
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"net/http"
	"strconv"
)

func errorHandler(res http.ResponseWriter, next *bool, msg string) {
	m, _ := json.Marshal(types.Success{Success: false, Msg: msg})
	methods.SendJSONResponse(res, m, 401)
	*next = false
}

func IsMaid(res http.ResponseWriter, req *http.Request) (next bool) {
	i, err := strconv.Atoi(req.Header.Get("Maid"))
	if err != nil || i == 0 {
		errorHandler(res, &next, "Not Maid")
	}

	next = true
	return
}
