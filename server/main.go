package main

import (
	"fmt"
	"log"
	"net/http"
	"video-manager/db"
	"video-manager/routers"
	"video-manager/utils"

	"github.com/gorilla/mux"
)

func buildRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routers.HomeRouter(router)
	return router
}

func main() {
	dbConnection := db.DatabaseConnection
	// Close the database connection before exiting the main function.
	defer dbConnection.Close()

	router := buildRouter()
	fmt.Println(fmt.Sprintf("Video Manager will listen on port %d...", utils.Config.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", utils.Config.Port), router); err != nil {
		log.Fatalf("could not listen on port %d %v", utils.Config.Port, err)
	}
}
