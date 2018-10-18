package methods

var validAddress = make(map[string]string)
var validPincode = make(map[string]string)

var pincodes []string
var locality []string

func init() {

	pincodes = []string{
		"560079",
	}

	locality = []string{
		"Basaveshwaranagar",
	}

	gos := 5

	proc := make(chan bool, gos)
	allDone := make(chan bool)

	for i := 0; i < gos; i++ {
		go func(n int) {
			for j := n; j < len(locality); j += gos {
				validAddress[pincodes[j]] = locality[j]
				validPincode[locality[j]] = pincodes[j]
			}
			proc <- true
		}(i)
	}

	go func() {
		for i := 0; i < gos; i++ {
			<-proc
		}
		allDone <- true
	}()

	for {
		select {
		case <-allDone:
			return
		}
	}

}

func IsPresent(pincode string, locality string) (check bool) {
	var temp string
	check = true
	temp, t := validAddress[pincode]

	if t && (temp == locality) {
		check = false
	}

	return
}
