package api

import (
	"encoding/json"
	"io"
)

type resp[T any] struct {
	StatusCode int `json:"statusCode"`
	Data       T   `json:"data,omitempty"`
}

func Json[T any](w io.Writer, statusCode int, data T) {
	json.NewEncoder(w).Encode(resp[T]{
		StatusCode: statusCode,
		Data:       data,
	})
}
