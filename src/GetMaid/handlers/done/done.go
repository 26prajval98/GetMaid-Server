package done

import (
	"GetMaid/database"
	"fmt"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) error {

	_ = res
	db := database.GetDb()

	id := req.URL.Query().Get("id")

	fmt.Println(req.URL.Query())

	_, e := db.Exec("UPDATE services SET DONE = 1 WHERE Service_id=?", id)

	fmt.Println(e)
	return nil
}
