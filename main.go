package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/routes"
	"github.com/chaseWilliams/family-map/lib/util"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

/*
TODO
- routes don't check for bad methods right now
*/

/*
main will set up the handlers and then start the server

NOTE
The test cases will drop the connection if an exception is thrown while receiving the
data from the server. Because of this, the handler may not go to completion.
Pretty piss poor implementation but whatever.
*/
func main() {
	http.HandleFunc("/", route)
	fmt.Println("serving at localhost:5000")
	http.ListenAndServe(":5000", nil)
}

/*
Sets a wrapper function to all service functions that goes and sets the appropriate headers
*/
func route(w http.ResponseWriter, r *http.Request) {
	var service routes.Route
	auth := false
	json := true

	// routing logic
	switch {
	case regexp.MustCompile(`/user/register`).MatchString(r.URL.Path):
		service = routes.Register
	case regexp.MustCompile(`/user/login`).MatchString(r.URL.Path):
		service = routes.Login
	case regexp.MustCompile(`/clear`).MatchString(r.URL.Path):
		service = routes.Clear
	case regexp.MustCompile(`/fill/\w*(/\d*)?`).MatchString(r.URL.Path):
		service = routes.FillUser
	case regexp.MustCompile(`/load`).MatchString(r.URL.Path):
		service = routes.Load
	case regexp.MustCompile(`/person/\w+`).MatchString(r.URL.Path):
		service = routes.GetPerson
		auth = true
	case regexp.MustCompile(`/person/?$`).MatchString(r.URL.Path):
		service = routes.GetFamily
		auth = true
	case regexp.MustCompile(`/event/\w+`).MatchString(r.URL.Path):
		service = routes.GetEvent
		auth = true
	case regexp.MustCompile(`/event/?$`).MatchString(r.URL.Path):
		service = routes.AllEvents
		auth = true
	default:
		service = routes.FileServer
		json = false
	}

	// initial setup
	if json {
		w.Header().Set("Content-Type", "application/json")
	}
	database.StartSession() // this sets the global database state within the scope of the request
	var err error = nil
	lrw := &loggingResponseWriter{w, http.StatusOK}
	defer deferredDatabaseCleanup(lrw, &err) // guaranteed database cleanup

	// auth and service logic
	var user models.User
	if auth {
		user, ok := authenticate(r)
		if !ok {
			util.WriteBadResponse(
				lrw,
				"invalid or missing authorization token",
			)
		} else {
			err = service(lrw, r, user)
		}
	} else {
		err = service(lrw, r, user)
	}

	// logging
	log.Printf("request at %s resulted in %v\n", r.URL.Path, lrw.StatusCode)
}

func authenticate(r *http.Request) (user models.User, ok bool) {
	token, ok := r.Header["Authorization"]
	if !ok {
		return
	}
	user, ok = models.AssertAuth(token[0])
	return
}

func deferredDatabaseCleanup(w http.ResponseWriter, err *error) {
	tx, _ := database.GetTransaction()
	// if panicking, rollback and escalate panic
	// else if service func returned an error, rollback
	if p := recover(); p != nil {
		tx.Rollback()
		log.Printf("PANIC: %v", p)
		util.WriteInternalServerError(w)
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
