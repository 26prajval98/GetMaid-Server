package methods

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(1000 + rand.Int()%9000)
}
