package response

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	statusCode int
	result     interface{}
}

func NewSuccess(result interface{}, status int) Success {
	return Success{
		result:     result,
		statusCode: status,
	}
}

func (success Success) Send(writer http.ResponseWriter) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(success.statusCode)
	return json.NewEncoder(writer).Encode(success.result)
}
