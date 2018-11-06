package methods

import (
	"GetMaid/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var Insert *sql.Stmt

type name struct {
	Name string `json:"name"`
}

func Handler(res http.ResponseWriter, req *http.Request) error {
	var e error
	defer ErrorHandler(res, &e)
	db := database.GetDb()

	//noinspection SqlResolve
	Insert, e := db.Prepare(`INSERT INTO image_upload  values (?,?)`)
	if e != nil {
		log.Fatal(e.Error())
	}

	maidid := req.Header.Get("Maid_id")
	RandString := randSeq(7)
	toret := maidid + RandString
	hpw, e := bcrypt.GenerateFromPassword([]byte(toret), 6)

	Insert.Exec(maidid, string(hpw))

	fmt.Println(string(hpw))

	jsonResp, _ := json.Marshal(name{Name: string(hpw)})
	SendJSONResponse(res, jsonResp, 200)

	return e
}
