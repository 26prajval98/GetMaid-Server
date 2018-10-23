package methods

import (
	"GetMaid/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Distance struct {
	Distance float64 `json:"distance"`
}

var (
	pincodeinsert *sql.Stmt
	UqPc          []string
)

func UniquePincodes(pincodes []string) []string {
	pc := make(map[string]bool)

	for _, entry := range pincodes {
		if _, value := pc[entry]; !value {
			pc[entry] = true
			UqPc = append(UqPc, entry)
		}
	}
	return UqPc
}

func FindDis(body []byte) Distance {
	var s Distance
	err := json.Unmarshal([]byte(body), &s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}

func initA() {

	Uniquepins := UniquePincodes(pincodes)

	var e error
	db := database.GetDb()

	var n, x int

	n = 0

	b := make(chan int)

	//noinspection SqlResolve
	pincodeinsert, e = db.Prepare(`INSERT INTO pincodes(Pincode1,Pincode2) VALUES ( ?, ?)`)

	for i, p1 := range Uniquepins {
		go func(n int, pin string) {

			defer func() {
				b <- 1
			}()

			for j := n + 1; j < len(Uniquepins); j++ {
				p2 := Uniquepins[j]
				w, err := http.Get("http://localhost:8000/distance/" + pin + "/" + p2)

				if err != nil {
					fmt.Println("Cannot get longitude and latitude for given pincode", pin, p2)
				}

				responseData, err := ioutil.ReadAll(w.Body)

				if err != nil {
					panic(err.Error())
				}
				s := FindDis(responseData)

				temp := s.Distance
				if temp < 3 {
					_, e = pincodeinsert.Exec(pin, p2)
				}
				if e != nil {
					panic(e.Error())
				}
			}
		}(i, p1)
	}

	for {
		select {
		case x = <-b:
			n += x
			if n == len(Uniquepins) {
				return
			}
		}
	}

	fmt.Println("Done")
}
