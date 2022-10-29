package api

import (
	"encoding/json"
	"net/http"
)

type resp[T any] struct {
	StatusCode int `json:"statusCode"`
	Data       T   `json:"data,omitempty"`
}

func Json[T any](w http.ResponseWriter, statusCode int, data T) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp[T]{
		StatusCode: statusCode,
		Data:       data,
	}) //#nosec
}
