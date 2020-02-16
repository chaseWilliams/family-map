package routes

import (
	"net/http"
	"encoding/json"
)

/* 
FillUser will fill out a user's ancestry tree, and remove previous ancestry data
if it exists.
*/
func FillUser(w http.ResponseWriter, r *http.Request) {
	m := map[string]string {"message": "success"}
	json.NewEncoder(w).Encode(m)
}