package util

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

/*
RandomID will generate a random ID that is 16 characters long
*/
func RandomID() string {
	lower := 65
	upper := 90
	bytes := make([]byte, 16)
	for i := range bytes {
		bytes[i] = byte(lower + rand.Intn(upper-lower+1))
	}
	return string(bytes)
}

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

/*
Message is a standard response struct with message and success
json fields
*/
type Message struct {
	Message string `json:"message"`
	Success string `json:"success"`
}

/*
WriteBadResponse will write a JSON error response object
with status code 400
*/
func WriteBadResponse(w http.ResponseWriter, message string) {
	WriteResponse(
		w,
		Message{
			Message: "error: " + message,
			Success: "false",
		},
		http.StatusBadRequest,
	)
}

/*
WriteOKResponse will dump the JSON as a response to the client
and respond with a 200
*/
func WriteOKResponse(w http.ResponseWriter, resp interface{}) {
	WriteResponse(
		w,
		resp,
		http.StatusOK,
	)
}

/*
WriteInternalServerError is a helper function that handles a generic 500 case
*/
func WriteInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "error: internal server error",
		"success": "false",
	})
}
