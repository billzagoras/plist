package http

import (
	"encoding/json"
	"net/http"
)

// Write json to response object.
func WriteJSON(status int, i interface{}, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(i)
}
