package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {

	var x int
	var err error

	t, _ := strconv.Atoi(os.Args[1])

	if t == 1 {
		fmt.Println("DB OPTION (0 :  Local) and (1 : db4free) NOTE : DEFAULT IS LOCAL ")
		fmt.Scanln(&x)
	} else {
		x = 0
	}

	if x != 1 {
		DB, err = sql.Open("mysql", "root@tcp(localhost:3306)/getmaid")
	} else {
		DB, err = sql.Open("mysql", "getmaid:mB79RWutLwFsS9v@tcp(db4free.net:3306)/getmaid")
	}

	if err != nil {
		log.Fatal("Database Not Connected")
	}

	createTables(DB)
}

func GetDb() *sql.DB {
	return DB
}

func CloseDb() {
	DB.Close()
}
