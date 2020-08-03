package db

import (
	"database/sql"
	"fmt"
	"video-manager/utils"
)

// var databaseConnection

// GetDatabaseConnection returns a connection to the dataset.
func GetDatabaseConnection() *sql.DB {
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
	return db
}
