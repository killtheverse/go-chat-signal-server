package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/killtheverse/go-chat-signal-server/config"
	"github.com/killtheverse/go-chat-signal-server/db"
	logger "github.com/killtheverse/go-chat-signal-server/logging"
	"go.mongodb.org/mongo-driver/mongo"
)

// App represents the web application that will be running on the server
type App struct {
    
    // The address on which the server will be started
    ServerAddress       string

    // Router for the app
    Router              *mux.Router

    // Database client for app
    DBClient            *mongo.Client

    // Database to access through client
    DB                  *mongo.Database
}

//ConfigurAppandRun creates a new instance of type App and runs it after configuring with the ServerConfig instance passed as argument
func ConfigurAppandRun(config *config.ServerConfig) {
    app := new(App)                
    app.initialize(config)
    app.run()
}

// initialize initializes an App instance and configure it according to the ServerConfig instance passed as argument
func (app *App) initialize(config *config.ServerConfig) {
    var err error
    app.ServerAddress = config.ServerAddress
    app.Router = mux.NewRouter()
    app.DBClient, err = db.Connect(config.DBURI)
    if err != nil {
        log.Fatal("[ERROR] Can't connect to database: %v", err)
    }
    app.DB = app.DBClient.Database(config.DBName)
    app.createIndexes()
    app.setupRouter()
}

// createIndexes creates unique indexes 
func (app *App) createIndexes() {

}

// setupRouter registers the routes 
func (app *App) setupRouter() {
    
}

// run starts the http server 
func (app *App) run () {
    server := http.Server{
        Addr: app.ServerAddress,
        Handler: app.Router,
        ReadTimeout: 5*time.Second,
        WriteTimeout: 10*time.Second,
        IdleTimeout: 120*time.Second,
    }

    go func() {
        logger.Write("Starting the server on: %v\n", app.ServerAddress)   
        err := server.ListenAndServe()
        if err != nil {
            logger.Fatal("[ERROR] Can't start the server: %v", err)
        }
    }()

    // Signals for shutting down the server
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

    // Block until a signal is received
    sig := <-sigs
    logger.Write("Trapped singal:%v\nShutting down the server\n", sig)

    // Disconnect the MongoDB client
    err := db.Disconnect(app.DBClient)
    if err != nil {
        logger.Write("[ERROR] Can't discconnect from database: %v\n", err)
    }

    // Shutdown the server, waiting for max 30 seconds
    logger.Write("Gracefully stopping server\n")
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    server.Shutdown(ctx)
}
