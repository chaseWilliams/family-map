package util

import (
	"net/http"
	"encoding/json"
)

/*
WriteResponse is a helper function that will write the body as JSON and set the
status code of the response object
*/
func WriteResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

/*
WriteNotFound is a helper function that handles a generic 404 case
TODO: replace with the 404 HTML page
*/
func WriteNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "error: resource not found",
		"success": "false",
	})
}