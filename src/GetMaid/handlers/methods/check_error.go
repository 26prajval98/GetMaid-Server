package methods

import "log"

func CheckErr(err error, str ...string) {

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	if err != nil {
		if len(str) == 0 {
			panic(err)
		} else {
			panic(str[0])
		}
	}
}
