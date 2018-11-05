package earnings

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"encoding/json"
	"net/http"
	"strconv"
)

type amount struct {
	Amount int `json:"earnings"`
}

func Handler(res http.ResponseWriter, req *http.Request) error {

	db := database.GetDb()

	mId, _ := strconv.Atoi(req.Header.Get("Maid_id"))

	rows, _ := db.Query("SELECT COUNT(*) as Worked FROM services WHERE Maid_id=? AND Done = 1", mId)

	rows.Next()

	var val int

	rows.Scan(&val)

	val = val * 80

	ja, _ := json.Marshal(amount{val})

	methods.SendJSONResponse(res, ja, 200)
	return nil
}
