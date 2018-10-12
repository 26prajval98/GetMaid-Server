package signup

import (
	"GetMaid/handlers/methods"
	"regexp"
)

func validateSignup(x interface{}) {
	var check bool
	var err error
	switch t := x.(type) {
	case Maid:
		if check, err = regexp.MatchString(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, t.Email); err == nil && !check {
			panic(EMAIL)
		} else if len(t.Name) == 0 {
			panic(NAME)
		} else if check, err = regexp.MatchString(`([1-9][0-9]+)`, t.Phone); len(t.Phone) != 10 && !check {
			panic(PHONE)
		} else if methods.IsPresent(t.Address.PinCode, t.Address.Locality) {
			panic(ADDRESS)
		}
	case Hirer:
		if check, err = regexp.MatchString(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, t.Email); err == nil && !check {
			panic(EMAIL)
		} else if len(t.Name) == 0 {
			panic(NAME)
		} else if check, err = regexp.MatchString(`([1-9][0-9]+)`, t.Phone); len(t.Phone) != 10 && !check {
			panic(PHONE)
		} else if methods.IsPresent(t.Address.PinCode, t.Address.Locality) {
			panic(ADDRESS)
		}
	}
}
