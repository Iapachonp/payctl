package database

import (

	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"

)



func Open() *sql.DB {
	db,error := sql.Open("sqlite3", "payctl.db")
	if error != nil {
		fmt.Println("payctl DB wasnt not able to start up: %v", error)
	}
	error = db.Ping()
	if error != nil {
		fmt.Println("DB is not avialable, not able to PING: %v", error)
	}
	// fmt.Println("Connected to payctl DB")
	return db
}
