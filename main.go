package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/chaseWilliams/family-map/lib/routes"
	"github.com/chaseWilliams/family-map/lib/database"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/jmoiron/sqlx"
)

/*
main will set up the handlers and then start the server
*/
func main() {
	setModelRoute("/user/login", routes.Login)
	fmt.Println("serving at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

/*
Sets a wrapper function to all service functions that goes and sets the appropriate headers
*/
func setModelRoute(path string, service func (w http.ResponseWriter, r *http.Request) error ) {
	addUniversalAttributes := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		database.StartSession() // this sets the global database state within the scope of the request
		var err error = nil
		lrw := &loggingResponseWriter{w, http.StatusOK}
		defer func() {
			tx := database.GetTransaction()
			// if panicking, rollback and escalate panic
			// else if service func returned an error, rollback
			if p := recover(); p != nil{
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
				log.Println(err)
			} else {
				tx.Commit()
			}
		}()
		err = service(lrw, r)
		log.Printf("request at %s resulted in %v\n", r.URL.Path, lrw.StatusCode)
	}
	http.HandleFunc(path, addUniversalAttributes)
}

type loggingResponseWriter struct {
    http.ResponseWriter
    StatusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
    lrw.StatusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}