package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/getmaid")

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
