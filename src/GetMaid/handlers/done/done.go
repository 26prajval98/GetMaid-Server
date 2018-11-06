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

	db.Exec("UPDATE services SET DONE = 1 WHERE Service_id=?", id)

	var maidId string

	x := db.QueryRow("SELECT Maid_id FROM services WHERE Service_id = ?", id)

	x.Scan(&maidId)

	t, err := db.Query("SELECT id FROM maid_online WHERE Maid_id= ?", maidId)

	if err != nil {
		fmt.Println(err.Error())
	}

	for t.Next() {
		var sid string
		t.Scan(&sid)
		db.Exec("UPDATE maid_online SET Count = Count - 1 WHERE id=?", sid)
	}

	return nil
}
