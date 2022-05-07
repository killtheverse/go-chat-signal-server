package util

import (
	"net/http"
	"time"

	"github.com/killtheverse/go-chat-signal-server/logging"
)

// JSONContentTypeMiddleware adds json content type header
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
        rw.Header().Add("content-type", "application/json; charset=UTF-8")
        next.ServeHTTP(rw, r)
    })
}

// LoggingMiddleware logs all incoming http requests
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(rw, r)
        logging.Write("%s %s %s", r.Method, r.RequestURI, time.Since(start))
    })
}
