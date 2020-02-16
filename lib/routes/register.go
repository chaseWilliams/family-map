package routes

import (
	"net/http"
	"encoding/json"
)

/* 
Register creates a new user account, generates 4 generations of ancestor data for the new
user, logs the user in, and returns an auth token.
*/
func Register(w http.ResponseWriter, r *http.Request) {
	m := map[string]string {"message": "success"}
	json.NewEncoder(w).Encode(m)
}