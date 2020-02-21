package util

import (
	"net/http"
	"encoding/json"
)

func WriteResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}