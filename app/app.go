package app

import (
	"net/http"

	"github.com/killtheverse/go-chat-signal-server/config"
)

// App is an interface which wraps the necessary modules for this system
type App interface {
    // Start all dependencies services
    Init() error

    // Start HTTP Server
    StartHttpServer() error
}

type restApiApplication struct {
    config      *config.Config
    httpServer  *http.Server
}


// Start other services
func (app *restApiApplication) Init() error {
    return nil
}

// Start and serve the HTTP Server
func(app *restApiApplication) StartHttpServer() error {
    return app.httpServer.ListenAndServe()
}
