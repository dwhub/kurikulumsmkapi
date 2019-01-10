package utils

import (
	"encoding/json"
	"net/http"
)

// Message - format response message
func Message(status int, message interface{}) map[string]interface{} {
	return map[string]interface{}{"code": status, "message": message}
}

// Respond - Update response
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.WriteHeader(data["code"].(int))
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
