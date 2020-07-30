package db

import (
	"database/sql"
)

// GetDatabaseConnection returns a connection to the dataset.
func GetDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database")
	if err != nil {
		panic(err.Error())
	}
	return db
}
