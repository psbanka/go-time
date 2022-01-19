package connector

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func sqlConnect() *sql.DB {
	db, err := sql.Open("mysql", "go-squee:my-new-password@(127.0.0.1:3306)/form_persistance?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
