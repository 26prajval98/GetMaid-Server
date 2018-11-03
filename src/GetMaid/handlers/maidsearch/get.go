package maidsearch

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
var (
	CountUpdate *sql.Stmt
	)



type maids struct {
	Id int `json:"id"`
}

func search(req *http.Request, res http.ResponseWriter) {
	var e error
	db := database.GetDb()
	maidsid := make([]maids, 0)
	q:=req.URL.Query()

	//reqServices:=req.URL.Query()["services"]
	reqServices := q.Get("services")
	hirerPincode := q.Get("pincode")
	//isactive:=q.Get("isActive")

	fmt.Println(reqServices,hirerPincode)

	if e != nil {
		log.Fatal(e.Error())
	}
	for Sid := range reqServices{
		CountUpdate,e= db.Prepare(`UPDATE maid SET Count=? WHERE Maid_id=?`)
		result, err := db.Query(`SELECT DISTINCT s.Maid_id,s.Count FROM maid s, maid_services x, address a, pincodes p WHERE x.Service_name=? AND s.AddressId=a.id AND ((a.Pincode=p.Pincode1 AND p.Pincode2=?) or a.Pincode = ?) AND x.Maid_id=s.Maid_id AND s.Active=1 AND s.Count<2`, Sid, hirerPincode, hirerPincode)


			log.Println(result)
			if err != nil {
				fmt.Println(result)
				log.Fatal(err.Error())
			}
			for result.Next() {
				var( id int
				ct int
				)
				err := result.Scan(&id,&ct)
				ct=ct+1
				CountUpdate.Exec(ct,id)
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
