package routes

import (
	"net/http"
	"encoding/json"
)

/* 
AllEvents Returns all events relevant to current user
*/
func AllEvents(w http.ResponseWriter, r *http.Request) {
	m := map[string]string {"message": "success"}
	json.NewEncoder(w).Encode(m)
}