package routes

import (
	"net/http"
	"encoding/json"
)

/* 
GetPerson will get the specified person object
*/
func GetPerson(w http.ResponseWriter, r *http.Request) {
	m := map[string]string {"message": "success"}
	json.NewEncoder(w).Encode(m)
}