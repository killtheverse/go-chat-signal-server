package util

import (
	"encoding/json"
	"net/http"
)

// Response to be returned to a http request
type response struct {
    // Status of the response
    Status      int         `json:"status"`

    // Message in the response
    Message     string      `json:"message"`

    // Content to be displayed
    Content     interface{} `json:"content"`
}

// ResponseWriter writes a http response to the response writer
func ResponseWriter(rw http.ResponseWriter, statusCode int, message string, content interface{}) error {
    rw.WriteHeader(statusCode)
    httpResponse := new(response)
    httpResponse.Status = statusCode
    httpResponse.Message = message
    httpResponse.Content = content
    err := json.NewEncoder(rw).Encode(httpResponse)
    return err
}
