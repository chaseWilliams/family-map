package routes

import (
	"net/http"
	"encoding/json"
)

/* 
GetFamily returns all related Persons to specified Person
*/
func GetFamily(w http.ResponseWriter, r *http.Request) {
	m := map[string]string {"message": "success"}
	json.NewEncoder(w).Encode(m)
}