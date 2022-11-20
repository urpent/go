package api

import (
	"encoding/json"
	"net/http"
)

type Status string

const (
	StatusOk           = "OK"
	StatusNotFound     = "NOT_FOUND"
	StatusBadRequest   = "BAD_REQUEST"
	StatusUnauthorized = "UNAUTHORIZED"
	StatusServerError  = "SERVER_ERROR"
)

type Response[T any] struct {
	Status  Status `json:"status"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

// List is used by query api in repo & service layer.
// Some clients require total records for pagination.
// Therefore, there will be an optional total field.
type List[T any] struct {
	Total  *int `json:"total,omitempty"`
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
	List   []T  `json:"list"`
}

func Json[T any](w http.ResponseWriter, status Status, data T) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response[T]{
		Status: status,
		Data:   data,
	}) //#nosec
}
