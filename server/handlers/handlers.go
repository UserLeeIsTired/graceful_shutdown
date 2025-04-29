package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteHeader(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
