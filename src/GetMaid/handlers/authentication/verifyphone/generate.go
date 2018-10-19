package verifyphone

import (
	"fmt"
	"github.com/subosito/twilio"
	"net/url"
)

const (
	accountSid = "ACbb0d964c185f8b35fc1c8b0b2eb4f78e"
	authToken  = "8f6ad6c609ff10518212535538b50ddc"
	phone      = "+18604075346"
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
