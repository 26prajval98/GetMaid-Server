package database

import (
	"database/sql"
	"log"
)

func checkErr(err error, str ...string) {

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

func createTables(db *sql.DB) {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS address (
  id int(11) NOT NULL AUTO_INCREMENT,
  Pincode varchar(6) NOT NULL,
  Locality varchar(100) NOT NULL,
  City varchar(100) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS hirer (
  Hirer_id int(11) NOT NULL AUTO_INCREMENT,
  Name varchar(100) NOT NULL,
  Email varchar(100) DEFAULT NULL,
  Phone varchar(10) DEFAULT NULL,
  HouseNumber varchar(5) DEFAULT NULL,
  AddressId int(11) DEFAULT NULL,
  Password varchar(1000) NOT NULL,
  Active int(1) NOT NULL,
  PRIMARY KEY (Hirer_id),
  UNIQUE KEY Email (Email),
  UNIQUE KEY Phone (Phone),
  CONSTRAINT fk_hirer FOREIGN KEY(AddressId) REFERENCES address(id) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS maid (
  Maid_id int(11) NOT NULL AUTO_INCREMENT,
  Name varchar(100) NOT NULL,
  Email varchar(100) DEFAULT NULL,
  Phone varchar(10) DEFAULT NULL,
  AddressId int(11) DEFAULT NULL,
  AccountId int(2) DEFAULT NULL,
  Password varchar(1000) NOT NULL,
  PRIMARY KEY (Maid_id),
  Active int(1) NOT NULL,
  UNIQUE KEY Email (Email),
  UNIQUE KEY Phone (Phone),
  CONSTRAINT fk_maid FOREIGN KEY(AddressId) REFERENCES address(id) ON DELETE SET NULL,
  CONSTRAINT fk_maid_ac FOREIGN KEY(AccountId) REFERENCES maid_card_details(acc_no) ON DELETE SET NULL

) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS pincodes(
  Pincode1 varchar(15) NOT NULL,
  Pincode2 varchar(15) NOT NULL,
  UNIQUE(Pincode1, Pincode2)
)ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS maid_work_timings(
	maid_id int(11) DEFAULT NULL,
	day enum("Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"),
	work_timings enum("00:00","01:00","02:00","03:00","04:00","05:00","06:00","07:00","08:00","09:00","10:00","11:00","12:00","13:00","14:00","15:00","16:00","17:00","18:00","19:00","20:00","21:00","22:00","23:00"),
	CONSTRAINT fk_maid_info FOREIGN KEY(maid_id) REFERENCES maid(Maid_id) ON DELETE SET NULL
)ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS maid_card_details(
	maid_id int(11) DEFAULT NULL,
	acc_no varchar(18) NOT NULL,
	ifsc_code varchar(11) ,
	CONSTRAINT fk_maid_card_details FOREIGN KEY(maid_id) REFERENCES maid(Maid_id) ON DELETE SET NULL
)ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS services (
  Service_id int(11) NOT NULL AUTO_INCREMENT,
  Maid_id int(11) DEFAULT NULL,
  Hirer_id int(11) DEFAULT NULL,
  Service_name enum("Cleaning", "Baby Sitting", "Washing Clothes", "Washing Utensils", "Gardening"),
  Done int(1) DEFAULT 0,
  Date DATETIME DEFAULT CURRENT_TIMESTAMP,
  Slot int(2),
  PRIMARY KEY(Service_id),
  CONSTRAINT fk_services_maid FOREIGN KEY(Maid_id) REFERENCES maid(Maid_id) ON DELETE SET NULL,
  CONSTRAINT fk_services_hirer FOREIGN KEY(Hirer_id) REFERENCES hirer(Hirer_id) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS service_time (
  Service_time_id int(11) NOT NULL AUTO_INCREMENT,
  Maid_id int(11) DEFAULT NULL,
  Slot int(2),
  Done int(1) DEFAULT 0,
  PRIMARY KEY(Service_time_id),
  CONSTRAINT fk_servicetime_maid FOREIGN KEY(Maid_id) REFERENCES maid(Maid_id) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS maid_services (
  Maid_services_id int(11) NOT NULL AUTO_INCREMENT,
  Maid_id int(11) DEFAULT NULL,
  Service_name enum("Cleaning", "Baby Sitting", "Washing Clothes", "Washing Utensils", "Gardening") DEFAULT "Cleaning" NOT NULL,
  PRIMARY KEY(Maid_services_id),
  UNIQUE(Maid_id, Service_name),
  CONSTRAINT fk_maidservice_maid FOREIGN KEY(Maid_id) REFERENCES maid(Maid_id) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;`,

		`CREATE TABLE IF NOT EXISTS maid_online (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  Maid_id int(11) DEFAULT NULL,
  Works int(1) DEFAULT 0,
  UNIQUE(Maid_id),
  CONSTRAINT fk_maidonline_maid FOREIGN KEY(Maid_id) REFERENCES maid(Maid_id) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;`,
	}

	gos := 5

	proc := make(chan bool)
	allDone := make(chan bool)

	for i := 0; i < gos; i++ {
		go func(n int) {
			for j := n; j < len(tables); j += gos {
				_, err := db.Exec(tables[j])

				if err != nil {
					checkErr(err, err.Error())
				}
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
