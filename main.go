package main

import (
	"fmt"
	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/routes"
	"github.com/chaseWilliams/family-map/lib/util"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

/*
main will set up the handlers and then start the server
*/
func main() {
	setModelRoute("/user/login", routes.Login)
	setModelRoute("/", routes.GetPerson) // all routes that don't match other route patterns
	fmt.Println("serving at localhost:5000")
	http.ListenAndServe(":5000", nil)
}

/*
Sets a wrapper function to all service functions that goes and sets the appropriate headers
*/
func setModelRoute(path string, method string, service func(w http.ResponseWriter, r *http.Request) error) {
	genericHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			util.WriteNotFound()
			log.Printf("request at %s was a %s instead of %s request", r.URL.Path, r.Method, method)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		database.StartSession() // this sets the global database state within the scope of the request
		var err error = nil
		lrw := &loggingResponseWriter{w, http.StatusOK}
		// deferred function
		defer deferredDatabaseCleanup(&err)
		err = service(lrw, r)
		log.Printf("request at %s resulted in %v\n", r.URL.Path, lrw.StatusCode)
	}
	http.HandleFunc(path, genericHandlerFunc)
}

func deferredDatabaseCleanup(err *error) {
	tx := database.GetTransaction()
	// if panicking, rollback and escalate panic
	// else if service func returned an error, rollback
	if p := recover(); p != nil {
		tx.Rollback()
		log.Printf("PANIC: %v", p)
		panic(p)
	} else if *err != nil {
		tx.Rollback()
		log.Println(*err)
	} else {
		tx.Commit()
	}
}

/*
loggingResponseWriter is an extended version of ResponseWriter that keeps track of the current
status code for logging purposes
*/
type loggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
