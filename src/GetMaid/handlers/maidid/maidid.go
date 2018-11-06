package maidid

import (
	"GetMaid/handlers/methods"
	"encoding/json"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) error {
	_ = req

	maidId := req.Header.Get("Maid_id")

	success, _ := json.Marshal(maidId)
	methods.SendJSONResponse(res, success, 200)
	return nil
}
