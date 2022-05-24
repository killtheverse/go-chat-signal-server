package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/killtheverse/go-chat-signal-server/internal/repositories"
)


func main() {
    mongoURI := os.Getenv("MONGO_URI")
    mongoDBName := os.Getenv("MONGO_DB_NAME")
    userHandler, err := InitializeUserHandlers(repositories.ConnectionURI(mongoURI), repositories.DatabaseName(mongoDBName))
    if err != nil {
        fmt.Print(err)
    }

    router := mux.NewRouter()
    router.HandleFunc("/users/register/", userHandler.Register).Methods("POST")

    server := http.Server{
        Addr: ":" + os.Getenv("PORT"),
        Handler: router,
        ReadTimeout: 5*time.Second,
        WriteTimeout: 10*time.Second,
        IdleTimeout: 120*time.Second,
    }

    log.Fatal(server.ListenAndServe())
}
