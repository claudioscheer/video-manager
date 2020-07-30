package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func buildRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home!")
	})
	return router
}

func main() {
	fmt.Println("Starting Video Manager server...")
	router := buildRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
