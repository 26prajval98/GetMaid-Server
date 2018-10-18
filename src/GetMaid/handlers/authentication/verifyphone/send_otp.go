package verifyphone

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func generateSecret(s []byte) (x string) {
	x = string(s)
	temp := strconv.Itoa(int(time.Now().Add(5*time.Minute).Unix() % 100000))
	x = temp + x
	return
}

func Send(res http.ResponseWriter, req *http.Request) {

	var err error
	var databaseActive int
	var isMaid int

	defer methods.ErrorHandler(res, &err)

	isMaid, err = strconv.Atoi(req.Header.Get("Maid"))

	db := database.GetDb()

	if isMaid == 0 {
		err := db.QueryRow("Select Active FROM hirer WHERE Phone=?", req.Header.Get("Phone")).Scan(&databaseActive)
		if err != nil {
			fmt.Println(err)
			panic(7)
			return
		}
	} else {
		err = db.QueryRow("SELECT Active FROM maid WHERE Phone=?", req.Header.Get("Phone")).Scan(&databaseActive)
		if err != nil {
			fmt.Println(err.Error())
			panic(7)
			return
		}
	}

	otp := methods.GenerateOTP()
	if !SendOTP(otp, req.Header.Get("Phone")) {
		panic(7)
	}

	temp, _ := bcrypt.GenerateFromPassword([]byte(otp), bcrypt.MinCost)

	otp = string(temp)

	success, err := json.Marshal(types.TokenUnverified{Success: true, Msg: req.Header.Get("Msg"), Secret: generateSecret(temp)})
	methods.CheckErr(err)
	methods.SendJSONResponse(res, success, 200)
}
