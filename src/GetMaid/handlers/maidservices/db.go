package maidservices

import (
	"GetMaid/database"
	"database/sql"
	"log"
)

var (
	serviceInsert *sql.Stmt
	serviceDelete *sql.Stmt
)

func init() {
	db := database.GetDb()

	var e error

	//noinspection SqlResolve
	serviceInsert, e = db.Prepare(`INSERT IGNORE INTO maid_services(Maid_id, Service_name) VALUES (?, ?)`)

	if e != nil {
		log.Fatal(e.Error())
	}

	//noinspection SqlResolve
	serviceDelete, e = db.Prepare(`DELETE FROM maid_services WHERE Maid_services_id=?`)

	if e != nil {
		log.Fatal(e.Error())
	}

}

func getServices(MaidId int) (bool, []service) {
	db := database.GetDb()
	ser := make([]service, 0)

	//noinspection SqlResolve
	rows, err := db.Query(`select Maid_services_id, Service_name from maid_services where Maid_id = ?`, MaidId)
	if err != nil {
		return false, ser
	}
	for rows.Next() {
		var id int
		var serviceName string
		err := rows.Scan(&id, &serviceName)

		if err != nil {
			return false, ser
		}

		ser = append(ser, service{id, serviceName})
	}
	return true, ser
}

func insertService(MaidId int, ServiceName string) bool {
	_, e := serviceInsert.Exec(MaidId, ServiceName)
	if e != nil {
		return false
	}
	return true
}

func deleteService(MaidServicesId int) bool {
	_, e := serviceDelete.Exec(MaidServicesId)
	if e != nil {
		return false
	}
	return true
}

//INSERT INTO table_listnames (name, address, tele)
//SELECT * FROM (SELECT 'Rupert', 'Somewhere', '022') AS tmp
//WHERE NOT EXISTS (
//SELECT name FROM table_listnames WHERE name = 'Rupert'
//) LIMIT 1;
