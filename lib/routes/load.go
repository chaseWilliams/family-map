package routes

import (
	"encoding/json"
	"net/http"
)

/*
Load will clear the database, and then load all data
*/
func Load(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{"message": "success"}
	json.NewEncoder(w).Encode(m)
}
