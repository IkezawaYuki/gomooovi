package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DbConnection *sql.DB

func init(){

	var err error
	DbConnection, err = sql.Open("mysql", "root")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("database is connected")
}