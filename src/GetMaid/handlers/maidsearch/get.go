package maidsearch

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"encoding/json"
	"log"
	"net/http"
)
type maids struct {
	Id int    `json:"id"`
}

func search(req *http.Request,res http.ResponseWriter){
	//var e error
	db := database.GetDb()
	maidsid := make([]maids, 0)
	reqServices:=req.URL.Query()["services"]
	hirerPincode:=req.URL.Query()["pincode"]
	//fmt.Println(services,"\n",pincode)
	for Sid:= range reqServices{
		result, err := db.Query(`SELECT maid_id FROM maid_services,address,pincodes WHERE Service_name=? AND  Pincode1=? AND Pincode2=Pincode`,Sid,hirerPincode)
		if err != nil {
			log.Fatal(err.Error())
		}
		for result.Next() {
			var id int
			err := result.Scan(&id)

			if err != nil {
				log.Fatal(err.Error())
			}

			maidsid = append(maidsid, maids{id})
		}

	}
	jsonResp, _ := json.Marshal(maidsid)

	methods.SendJSONResponse(res, jsonResp, 200)



	//req.URL.Query()
	//log.Println(req.URL.Query())
	//success, err := json.Marshal(types.Success{Success: true, Msg: "Hello"})
	//methods.CheckErr(err)
	//methods.SendJSONResponse(res, success, 400)
}
