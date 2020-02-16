package main

import (
	"log"
	"net/http"
	"github.com/chaseWilliams/family-map/lib/routes"
)

/*
main will set up the handlers and then start the server
*/
func main() {
	setModelRoute("/", routes.Register)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

/*
Sets a wrapper function to all service functions that goes and sets the appropriate headers
*/
func setModelRoute(path string, service func (w http.ResponseWriter, r *http.Request) ) {
	addUniversalAttributes := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		service(w, r)
	}
	http.HandleFunc(path, addUniversalAttributes)
}