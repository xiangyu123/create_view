package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	// SqlDB, err = sql.Open("mysql", "root:InnoTree20%217%40nx6l8@tcp(172.31.215.34:3306)/zzz?charset=utf-8")
	SqlDB, err = sql.Open("mysql", "admin_m:InnoTree20!7@nx6l8@tcp(172.31.215.34:3306)/zzz")
	if err != nil {
		log.Fatal(err.Error())
	}

	SqlDB.SetMaxIdleConns(3)
	SqlDB.SetMaxOpenConns(3)

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
