package maidservices

import (
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"net/http"
	"strconv"
)

func post(res http.ResponseWriter, Id, Service string) {
	var err error
	defer methods.ErrorHandler(res, &err)

	MaidId, _ := strconv.Atoi(Id)

	b := insertService(MaidId, Service)
	if !b {
		panic(7)
	}

	success, _ := json.Marshal(types.Success{Success: true})

	methods.SendJSONResponse(res, success, 200)
}
