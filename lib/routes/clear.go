package routes

import (
	"encoding/json"
	"net/http"
)

/*
Clear will wipe all data from the database
*/
func Clear(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{"message": "success"}
	json.NewEncoder(w).Encode(m)
}
