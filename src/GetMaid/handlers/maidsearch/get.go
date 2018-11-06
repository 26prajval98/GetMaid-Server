package maidsearch

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	CountUpdate *sql.Stmt
)

type maids struct {
	Id      int    `json:"id"`
	Service string `json:"service"`
}

func search(req *http.Request, res http.ResponseWriter) {
	var e error
	var maidsid maids

	db := database.GetDb()
	q := req.URL.Query()

	reqServices := q.Get("services")
	hirerPincode := q.Get("pincode")

	if e != nil {
		log.Fatal(e.Error())
	}

	//noinspection SqlResolve
	CountUpdate, e = db.Prepare(`UPDATE maid_online SET Count=Count+1 WHERE Maid_id=?`)

	if e != nil {
		fmt.Println(e.Error())
	}

	//noinspection SqlResolve
	result, err := db.Query(`SELECT DISTINCT o.Maid_id, o.Count FROM maid s, maid_services x, address a, pincodes p, maid_online o WHERE x.Service_name=? AND s.AddressId=a.id AND ((a.Pincode=p.Pincode1 AND p.Pincode2=?) or a.Pincode = ?) AND x.Maid_id=s.Maid_id AND s.Maid_id=o.Maid_id AND o.Count<2 LIMIT  1`, reqServices, hirerPincode, hirerPincode)

	if err != nil {
		log.Fatal(err.Error())
	}

	var (
		id int
		ct int
	)

	id = -1

	for result.Next() {
		err := result.Scan(&id, &ct)
		CountUpdate.Exec(id)
		if err != nil {
			log.Fatal(err.Error())
		}

		maidsid = maids{id, reqServices}
	}

	hirerId, _ := strconv.Atoi(req.Header.Get("Hirer_id"))

	if id != -1 {
		//noinspection SqlResolve
		db.Exec("INSERT INTO services(Maid_id, Hirer_id, Service_name) VALUES (?, ?, ?)", id, hirerId, reqServices)
		jsonResp, _ := json.Marshal(maidsid)
		methods.SendJSONResponse(res, jsonResp, 200)
	} else {
		jsonResp, _ := json.Marshal(types.Success{Success: false, Msg: "No Maids Found Try Again Later"})
		methods.SendJSONResponse(res, jsonResp, 200)
	}
}
