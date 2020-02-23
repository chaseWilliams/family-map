package routes

import (
	"encoding/json"
	"net/http"
)

/*
GetEvent returns specified event object
*/
func GetEvent(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{"message": "success"}
	json.NewEncoder(w).Encode(m)
}
