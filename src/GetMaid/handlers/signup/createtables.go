package signup

import (
	"GetMaid/handlers/methods"
	"database/sql"
)

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
  Password varchar(1000) NOT NULL,
  PRIMARY KEY (Maid_id),
  UNIQUE KEY Phone (Phone),
  CONSTRAINT fk_maid FOREIGN KEY(AddressId) REFERENCES address(id) ON DELETE SET NULL
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
					methods.CheckErr(err, err.Error())
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
