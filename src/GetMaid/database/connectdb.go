package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "mB79RWutLwFsS9v@tcp(db4free.net:3306)/getmaid")

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