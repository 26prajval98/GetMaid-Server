package maidservices

import (
	"GetMaid/handlers/methods"
	"encoding/json"
	"net/http"
	"strconv"
)

type service struct {
	Id          int    `json:"id"`
	ServiceName string `json:"service_name"`
}

type serviceArray struct {
	Services []service `json:"services"`
}

func get(res http.ResponseWriter, msg string) {
	var err error
	defer methods.ErrorHandler(res, &err)
	MaidId, _ := strconv.Atoi(msg)
	b, s := getServices(MaidId)
	if !b {
		panic(7)
	}

	serviceJson, _ := json.Marshal(serviceArray{s})

	methods.SendJSONResponse(res, serviceJson, 200)
}
