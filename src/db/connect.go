package db

import (
	"database/sql"
	"fmt"
	"video-manager/utils"
)

var databaseConnection *sql.DB = nil

// GetDatabaseConnection returns a connection to the dataset.
func GetDatabaseConnection() *sql.DB {
	if databaseConnection != nil {
		return databaseConnection
	}
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			utils.Config.Database.User,
			utils.Config.Database.Password,
			utils.Config.Database.Host,
			utils.Config.Database.Port,
			utils.Config.Database.DB,
		),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	databaseConnection = db
	return databaseConnection
}
