package signup

import (
	"GetMaid/database"
	"GetMaid/handlers/methods"
	"database/sql"
	"errors"
)

type d struct {
	error bool
	e     error
}

var (
	maidInsert    *sql.Stmt
	hirerInsert   *sql.Stmt
	addressInsert *sql.Stmt
	addressDelete *sql.Stmt
)

func init() {
	var e error
	db := database.GetDb()

	createTables(db)

	//noinspection SqlResolve
	maidInsert, e = db.Prepare(`INSERT INTO maid(Name, Email, Phone, AddressId, Password, Active) VALUES ( ?, ?, ?, ?, ?, 0)`)
	methods.CheckErr(e)

	//noinspection SqlResolve
	hirerInsert, e = db.Prepare(`INSERT INTO hirer(Name, Email, Phone, HouseNumber, AddressId, Password, Active) VALUES ( ?, ?, ?, ?, ?, ?, 0)`)
	methods.CheckErr(e)

	//noinspection SqlResolve
	addressInsert, e = db.Prepare(`INSERT INTO address(Pincode, Locality, City) VALUES ( ?, ?, ?)`)
	methods.CheckErr(e)

	//noinspection SqlResolve
	addressDelete, e = db.Prepare(`DELETE FROM address WHERE id=?`)
	methods.CheckErr(e)
}

func insertMaid(user Maid) d {

	r := make(chan sql.Result)
	done := make(chan d)
	var temp d

	go func() {

		defer func() {
			if r := recover(); r != nil {
				switch str := r.(type) {
				case string:
					temp = d{true, errors.New(str)}
				}
			}
		}()

		re, er := addressInsert.Exec(user.Address.PinCode, user.Address.Locality, user.Address.City)
		if er != nil {
			done <- d{true, er}
			panic(er.Error())
		}
		r <- re
	}()

	go func() {

		defer func() {
			if r := recover(); r != nil {
				switch str := r.(type) {
				case string:
					temp = d{true, errors.New(str)}
				}
			}
		}()

		var id int64
		var er error
		var email interface{}
		t := <-r
		id, er = t.LastInsertId()
		if er != nil {
			done <- d{true, er}
			panic(er.Error())
		}

		if user.Email == "" {
			email = nil
		} else {
			email = user.Email
		}
		_, er = maidInsert.Exec(user.Name, email, user.Phone, id, user.Password)
		if er != nil {

			addressDelete.Exec(id)

			done <- d{true, errors.New("email / phone already exists")}
		}
		done <- d{false, nil}
	}()

	for {
		select {
		case temp = <-done:
			return temp
		}
	}
}

func insertHirer(user Hirer) d {

	r := make(chan sql.Result)
	done := make(chan d)
	var temp d

	go func() {

		defer func() {
			if r := recover(); r != nil {
				switch str := r.(type) {
				case string:
					temp = d{true, errors.New(str)}
				}
			}
		}()

		re, er := addressInsert.Exec(user.Address.PinCode, user.Address.Locality, user.Address.City)
		if er != nil {
			done <- d{true, er}
			panic(er.Error())
		}
		r <- re
	}()

	go func() {

		defer func() {
			if r := recover(); r != nil {
				switch str := r.(type) {
				case string:
					temp = d{true, errors.New(str)}
				}
			}
		}()

		var id int64
		var er error
		t := <-r
		id, er = t.LastInsertId()
		if er != nil {
			done <- d{true, er}
			panic(er.Error())
		}
		_, er = hirerInsert.Exec(user.Name, user.Email, user.Phone, user.HouseNo, id, user.Password)
		if er != nil {

			addressDelete.Exec(id)

			done <- d{true, errors.New("email / phone already exists")}
		}
		done <- d{false, nil}
	}()

	for {
		select {
		case temp = <-done:
			return temp
		}
	}
}
