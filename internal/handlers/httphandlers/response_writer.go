package httphandlers

import (
	"encoding/json"
	"net/http"

	"github.com/killtheverse/go-chat-signal-server/internal/handlers/httphandlers/response"
)

// WriteResponse writes an http response
func WriteResponse(rw http.ResponseWriter, statusCode int, message string, content interface{}) error {
    rw.WriteHeader(statusCode)
    httpResponse := new(response.Response)
    httpResponse.Status = statusCode
    httpResponse.Message = message
    httpResponse.Content = content
    err := json.NewEncoder(rw).Encode(httpResponse)
    return err
}
