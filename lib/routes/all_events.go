package routes

import (
	"encoding/json"
	"net/http"
)

/*
AllEvents Returns all events relevant to current user
*/
func AllEvents(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{"message": "success"}
	json.NewEncoder(w).Encode(m)
}
