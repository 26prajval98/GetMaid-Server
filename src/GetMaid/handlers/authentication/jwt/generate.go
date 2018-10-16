package jwt

import (
	"github.com/gbrlsnchs/jwt"
	"time"
)

func GenerateJWT(id int64, phone string, maid bool) (token []byte, err error) {

	var payload []byte
	now := time.Now()

	hs256 := jwt.NewHS256(key)

	jot := &jwt.JWT{
		Issuer:         "GetMaid",
		Subject:        "Auth",
		Audience:       "User",
		ExpirationTime: now.Add(12 * time.Hour).Unix(),
		IssuedAt:       now.Unix(),
		ID:             id,
		Phone:          phone,
		Maid:           maid,
	}

	jot.SetAlgorithm(hs256)
	payload, err = jwt.Marshal(jot)

	if err != nil {
		return []byte("Something went wrong"), err
	}

	token, err = hs256.Sign(payload)

	if err != nil {
		return []byte("Something went wrong"), err
	}

	return
}
