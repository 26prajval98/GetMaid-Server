package maidservices

import (
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"net/http"
	"strconv"
)

func del(res http.ResponseWriter, MaidServiceIdS string) {
	var err error
	defer methods.ErrorHandler(res, &err)

	MaidServiceId, _ := strconv.Atoi(MaidServiceIdS)

	b := deleteService(MaidServiceId)

	if !b {
		panic(7)
	}

	success, _ := json.Marshal(types.Success{Success: true})

	methods.SendJSONResponse(res, success, 200)
}
