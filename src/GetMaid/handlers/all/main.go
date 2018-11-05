package pendingall

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"encoding/json"
	"fmt"
	"net/http"
)

type hirerPending struct {
	Sid string `json:"sid"`
	Name string `json:"Name"`
	Phone string `json:"Phone"`
	Work string `json:"work"`
}

func Handler(res http.ResponseWriter, req *http.Request) error {
	var e error
	defer methods.ErrorHandler(res, &e)
	_ = req

	var (
		hId   string
		mId string
		sname string
	)


	p := make([]hirerPending, 0)

	db := database.GetDb()

	hId = req.Header.Get("Hirer_id")

	rows, e := db.Query("SELECT Service_id, Maid_id, Service_name FROM services WHERE Hirer_id = ? AND Done=0", hId)

	if e != nil {
		fmt.Println(e.Error())
	}

	var (
		Name, Phone, sId string
	)

	for rows.Next() {
		rows.Scan(&sId,&mId, &sname)
		row := db.QueryRow("Select h.Name, h.Phone FROM maid h WHERE h.Maid_id=?", mId)
		row.Scan(&Name, &Phone)
		p = append(p, hirerPending{Sid : sId, Name:Name, Phone:Phone, Work:sname})
	}

	jp, _ := json.Marshal(p)

	methods.SendJSONResponse(res, jp, 200)

	return nil
}
