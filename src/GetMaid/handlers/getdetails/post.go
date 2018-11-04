package getdetails

import (
	"GetMaid/database"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func update(req *http.Request) {
	db := database.GetDb()

	IsMaid, _ := strconv.Atoi(req.Header.Get("Maid"))

	req.ParseForm()
	var (
		Aid *sql.Rows
		aId int
		e   error
	)
	if IsMaid == 1 {
		Name, Email, Phone := req.FormValue("Name"), req.FormValue("Email"), req.FormValue("Phone")

		id, _ := strconv.Atoi(req.Header.Get("Maid_id"))

		wg.Add(1)
		//noinspection SqlResolve
		db.Exec("UPDATE maid SET Name=?, Email=?, Phone=?, Active=0 WHERE Maid_id=?", Name, Email, Phone, id)

		fmt.Println(id)
		//noinspection SqlResolve
		Aid, e = db.Query("SELECT AddressId FROM maid WHERE Maid_id=?", id)

		Aid.Next()
		Aid.Scan(&aId)

	} else {
		Name, Email, Phone := req.FormValue("Name"), req.FormValue("Email"), req.FormValue("Phone")

		id, _ := strconv.Atoi(req.Header.Get("Hirer_id"))

		//noinspection SqlResolve
		db.Exec("UPDATE maid SET Name=?, Email=?, Phone=?, Active=0 WHERE Hirer_id=?", Name, Email, Phone, id)

		//noinspection SqlResolve
		Aid, _ = db.Query("SELECT AddressId from hirer WHERE Hirer_id=?", id)
		Aid.Next()
		Aid.Scan(&aId)
	}

	//noinspection SqlResolve
	db.Exec("UPDATE Address SET Pincode=?, Locality=? WHERE id=?", req.FormValue("Pincode"), req.FormValue("addr"), aId)

	_ = e
	return
}

func post(req *http.Request) {
	update(req)
	return
}
