package ismaid

import (
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"net/http"
	"strconv"
)

func Handler(res http.ResponseWriter, req *http.Request) error {
	_ = req

	var maid string

	ismaid, _ := strconv.Atoi(req.Header.Get("Maid"))

	if ismaid == 1 {
		maid = "MAID"
	} else {
		maid = "HIRER"
	}
	success, _ := json.Marshal(types.Success{Success: true, Msg: maid})
	methods.SendJSONResponse(res, success, 200)
	return nil
}
