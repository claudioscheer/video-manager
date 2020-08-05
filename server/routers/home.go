package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeRouter handle "/".
func HomeRouter(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	})
}
