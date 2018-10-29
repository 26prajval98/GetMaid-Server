package verifyphone

import (
	"fmt"
	"github.com/subosito/twilio"
	"net/url"
)

const (
	accountSid = "AC1c044bc7f4f75c9731ddf1b3934cf969"
	authToken  = "9a8e66793b7b95234248f02ab09a0fbd"
	phone      = "+19794815817"
)

func client(toPhone string, msg string) error {
	c := twilio.NewClient(accountSid, authToken, nil)

	s, _, err := c.Messages.Create(url.Values{
		"From": {phone},
		"To":   {"+91" + toPhone},
		"Body": {msg},
	})

	fmt.Println(s)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func SendOTP(otp, toPhone string) bool {
	msg := "Your OTP for GetMaid is " + otp
	err := client(toPhone, msg)

	if err != nil {
		return false
	}

	return true
}
