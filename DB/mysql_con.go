package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/pemilihan")
	if err != nil {
		panic(err)
	}
	DB = db
}
