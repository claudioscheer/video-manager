package db

import (
	"database/sql"
	"fmt"
	"video-manager/utils"

	// Import mysql driver.
	_ "github.com/go-sql-driver/mysql"
)

// DatabaseConnection stores the connection to the database. This works like a singleton class.
var DatabaseConnection *sql.DB = nil

func init() {
	dbConnection, err := getDatabaseConnection()
	if err != nil {
		panic(err)
	}
	if err = dbConnection.Ping(); err != nil {
		panic(err)
	}
	DatabaseConnection = dbConnection
}

func getDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?timeout=%ds",
			utils.Config.Database.User,
			utils.Config.Database.Password,
			utils.Config.Database.Host,
			utils.Config.Database.Port,
			utils.Config.Database.DB,
			utils.Config.Database.Timeout,
		),
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
