package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func init() {
	//db, err := sql.Open("mysql", "getmaid:mB79RWutLwFsS9v@tcp(db4free.net:3306)/getmaid")
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/getmaid")

	if err != nil {
		log.Fatal("Database Not Connected")
	}

	DB = db
}
func GetDb() *sql.DB {
	return DB
}

func CloseDb() {
	DB.Close()
}
