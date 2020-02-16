package routes

import (
	"net/http"
	"encoding/json"
)

/* 
GetEvent returns specified event object
*/
func GetEvent(w http.ResponseWriter, r *http.Request) {
	m := map[string]string {"message": "success"}
	json.NewEncoder(w).Encode(m)
}