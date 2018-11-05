package methods

import (
	"GetMaid/database"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

var Insert *sql.Stmt

func Handler(res http.ResponseWriter, req *http.Request) error {
	var e error
	defer ErrorHandler(res, &e)
	db := database.GetDb()

	//noinspection SqlResolve
	Insert, e := db.Prepare(`INSERT INTO image_upload values (?,?)`)
	if e != nil {
		log.Fatal(e.Error())
	}
	maidid := req.Header.Get("Maid_id")
	RandString := randSeq(7)
	mid := sha256.Sum256([]byte(maidid))
	m := string(mid[:])
	toret := m + RandString
	Insert.Exec(maidid, toret)
	jsonResp, _ := json.Marshal(toret)

	SendJSONResponse(res, jsonResp, 200)

	return e
}
