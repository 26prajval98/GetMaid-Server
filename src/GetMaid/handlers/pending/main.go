package pending

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"encoding/json"
	"fmt"
	"net/http"
)

type pending struct {
	Location string `json:"location"`
	Service  string `json:"service"`
}

func Handler(res http.ResponseWriter, req *http.Request) error {
	var e error
	defer methods.ErrorHandler(res, &e)
	_ = req

	p := make([]pending, 0)

	db := database.GetDb()

	mId := req.Header.Get("Maid_id")

	rows, e := db.Query("SELECT Hirer_id, Service_name FROM services WHERE Maid_id = ? AND Done=0", mId)

	if e != nil {
		fmt.Println(e.Error())
	}

	var (
		hId   int
		sname string
	)

	for rows.Next() {
		rows.Scan(&hId, &sname)
		var hno, pin, loc string
		row := db.QueryRow("Select h.HouseNumber, a.Pincode, a.Locality FROM hirer h, address a WHERE h.Hirer_id=? AND h.AddressId=a.id", hId)
		row.Scan(&hno, &pin, &loc)
		p = append(p, pending{Location: hno + ", " + loc + ", " + pin, Service: sname})
	}

	jp, _ := json.Marshal(p)

	methods.SendJSONResponse(res, jp, 200)
	return nil
}
