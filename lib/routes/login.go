package routes

import (
	"net/http"
	"encoding/json"
)

/* 
Login will login the user and return an auth token
*/
func Login(w http.ResponseWriter, r *http.Request) {
	m := map[string]string {"message": "success"}
	json.NewEncoder(w).Encode(m)
}