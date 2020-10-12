package server

import (
	"encoding/json"
	"log"
	"net/http"
)

//go:generate mockery --dir . --name ResponseWriter --case underscore --output mocks/
type ResponseWriter interface {
	http.ResponseWriter
}

// ResponseErr holds the error response
type ResponseErr struct {
	Error string `json:"error,omitempty"`
}

// SendHTTPResponse writes HTTP response back including HTTP code number and encode payload
func SendHTTPResponse(w ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		log.Printf("Warning: failed to write response [%v] to the response writer\n", w)
	}
}
