package getdetails

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Handler(res http.ResponseWriter, req *http.Request) error {
	var e error
	defer methods.ErrorHandler(res, &e)
	_ = req

	var rows *sql.Row
	ismaid, _ := strconv.Atoi(req.Header.Get("Maid"))

	db := database.GetDb()

	if ismaid == 1 {
		//noinspection SqlResolve
		rows = db.QueryRow("SELECT m.Name, m.Email, m.Phone, a.Pincode FROM maid m, address a WHERE a.id=m.AddressId AND m.Maid_id=?", req.Header.Get("Maid_id"))
		if e != nil {
			fmt.Println(e.Error())
			panic("Shit Happens")
		}
	} else {
		//noinspection SqlResolve
		rows = db.QueryRow("SELECT m.Name, m.Email, m.Phone, a.Pincode FROM hirer m, address a WHERE (a.id=m.AddressId AND m.Hirer_id=?)", req.Header.Get("Hirer_id"))
		if e != nil {
			fmt.Println(e.Error())
			panic("Shit Happens")
		}
	}

	var Name, Email, Phone, Pincode string

	rows.Scan(&Name, &Email, &Phone, &Pincode)

	success, _ := json.Marshal(types.Details{Name: Name, Email: Email, Phone: Phone, Pincode: Pincode})
	methods.SendJSONResponse(res, success, 200)
	return nil
}
