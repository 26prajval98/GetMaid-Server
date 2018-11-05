package verifyphone

import (
	"fmt"
	"github.com/subosito/twilio"
	"net/url"
)

const (
	accountSid = "AC2eac4de556bc49880e8d56c230713c73"
	authToken  = "6b2ad5c1f0da177c0730b7da4e342a04"
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
	//err := client(toPhone, msg)

	err := error(nil)

	_ = client

	fmt.Println(msg)

	if err != nil {
		return false
	}

	return true
}
