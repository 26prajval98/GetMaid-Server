package methods

import (
	"GetMaid/database"
	"encoding/json"
	"log"
	"net/http"
	"database/sql"
)
var Insert *sql.Stmt

type maid struct {

	Name string `json:"name"`
}
func Handler(res http.ResponseWriter,req *http.Request)error{
	var e error
	defer ErrorHandler(res, &e)
	db:=database.GetDb()
	Insert,e:= db.Prepare(`INSERT INTO image_upload values (?,?)`)
	if e!=nil{
		log.Fatal(e.Error())
	}
	maidid:=req.Header.Get("Maid_id")

	//Instead of random string name of the maid is being returned
	q,err:=db.Query(`SELECT m.Name from maid m WHERE m.Maid_id=?`,maidid)
	if err!=nil{
		log.Fatal(err.Error())
	}

	var name string
	err = q.Scan(&name)
	if err!=nil{
		log.Fatal(err.Error())
	}
	v:=maid{name}
	Insert.Exec(maidid,name)
	jsonResp, _ := json.Marshal(v)

	SendJSONResponse(res, jsonResp, 200)

	return e
}
