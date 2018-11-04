package maidonline

import (
	"GetMaid/database"
	"net/http"
	"strconv"
)

func get(req *http.Request) {
	db := database.GetDb()

	id, _ := strconv.Atoi(req.Header.Get("Maid_id"))

	db.Exec("INSERT INTO maid_online(Maid_id) VALUES(?)", id)
}

func getAdd(req *http.Request) {
	db := database.GetDb()

	id, _ := strconv.Atoi(req.Header.Get("Maid_id"))

	db.Exec("UPDATE maid_online SET Count = Count + 1 WHERE Maid_id=?", id)
}

func del(req *http.Request) {
	db := database.GetDb()

	id, _ := strconv.Atoi(req.Header.Get("Maid_id"))

	db.Exec("DELETE FROM maid_online WHERE Maid_id=?", id)
}
