package db

import (
	"database/sql"
)

func GetDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database")
	if err != nil {
		panic(err.Error())
	}
	return db
}
