package routes

import (
	"net/http"
	"encoding/json"
)

/* 
Clear will wipe all data from the database
*/
func Clear(w http.ResponseWriter, r *http.Request) {
	m := map[string]string {"message": "success"}
	json.NewEncoder(w).Encode(m)
}