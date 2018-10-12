package methods

var validAddress = make(map[string]string)

func init() {
	validAddress["560079"] = "Basaveshwaranagar"
}

func IsPresent(pincode string, locality string) (check bool) {
	var temp string
	check = true
	temp, check = validAddress[pincode]

	if check && temp == locality {
		check = false
	}
	return
}
