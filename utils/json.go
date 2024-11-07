package utils

import (
	"encoding/json"
	"golang-beginner-chap24/collections"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(collections.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}
