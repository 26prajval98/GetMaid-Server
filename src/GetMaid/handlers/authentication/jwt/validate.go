package jwt

import (
	"GetMaid/handlers/methods"
	"GetMaid/handlers/types"
	"encoding/json"
	"fmt"
	"github.com/gbrlsnchs/jwt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func errorHandler(res http.ResponseWriter, next *bool) {
	m, _ := json.Marshal(types.Success{Success: false, Msg: "Not Authenticated"})
	methods.SendJSONResponse(res, m, 401)
	*next = false
}

func VerifyJWT(res http.ResponseWriter, req *http.Request) (next bool) {
	crude := req.Header.Get("Authorization")

	tokens := strings.Split(crude, " ")

	if len(tokens) < 2 {
		errorHandler(res, &next)
		return
	}

	token := tokens[1]

	now := time.Now()

	hs256 := jwt.NewHS256(key)

	payload, sig, err := jwt.Parse(token)
	if err != nil {
		errorHandler(res, &next)
	}

	if err = hs256.Verify(payload, sig); err != nil {
		fmt.Println(err.Error())
		return
	}

	var jot jwt.JWT
	if err = jwt.Unmarshal(payload, &jot); err != nil {
		errorHandler(res, &next)
		return
	}

	iatValidator := jwt.IssuedAtValidator(now)
	expValidator := jwt.ExpirationTimeValidator(now)
	if err = jot.Validate(iatValidator, expValidator); err != nil {
		switch err {
		case jwt.ErrIatValidation:
			errorHandler(res, &next)
			return
		case jwt.ErrExpValidation:
			errorHandler(res, &next)
			return
		}
	}

	if jot.Maid {
		req.Header.Add("Maid_id", strconv.Itoa(int(jot.ID)))
		req.Header.Add("Phone", jot.Phone)
	} else {
		req.Header.Add("Hirer_id", strconv.Itoa(int(jot.ID)))
		req.Header.Add("Phone", jot.Phone)
	}

	next = true
	return
}
