package main

import (
	"fmt"
	"log"
	"net/http"
	"video-manager/db"
	"video-manager/utils"

	"github.com/gorilla/mux"
)

func buildRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	})
	return router
}

func main() {
	dbConnection := db.DatabaseConnection
	// Close the database connection before exiting the main function.
	defer dbConnection.Close()

	router := buildRouter()
	fmt.Println(fmt.Sprintf("Video Manager listen on port %d...", utils.Config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", utils.Config.Port), router))
}
