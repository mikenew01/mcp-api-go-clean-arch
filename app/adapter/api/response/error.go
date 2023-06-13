package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	statusCode int
	Errors     []string `json:"errors"`
}

func NewError(err error, status int) *Error {
	return &Error{
		Errors:     []string{err.Error()},
		statusCode: status,
	}
}

func (err Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.statusCode)
	return json.NewEncoder(w).Encode(err)
}
