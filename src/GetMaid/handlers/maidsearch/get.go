package maidsearch

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type maids struct {
	Id int `json:"id"`
}

func search(req *http.Request, res http.ResponseWriter) {
	//var e error
	db := database.GetDb()
	maidsid := make([]maids, 0)
	reqServices := req.URL.Query()["services"]
	hirerPincode := req.URL.Query()["pincode"]
	//fmt.Println(services,"\n",pincode)
	//SELECT DISTINCT s.Maid_id FROM maid s, maid_services x, address a, pincodes p WHERE s.AddressId=a.id AND ((a.Pincode=p.Pincode1 AND p.Pincode2="560025") or a.Pincode = "560025") AND x.Maid_id=s.Maid_id
	for Sid := range reqServices {
		result, err := db.Query(`SELECT DISTINCT s.Maid_id FROM maid s, maid_services x, address a, pincodes p WHERE s.AddressId=a.id AND a.Pincode=p.Pincode1 AND x.Maid_id=s.Maid_id`, Sid, hirerPincode)
		if err != nil {
			fmt.Println(result)
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
