package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/killtheverse/go-chat-signal-server/logging"
	"github.com/killtheverse/go-chat-signal-server/models"
	"github.com/killtheverse/go-chat-signal-server/util"
)

// Register creates a new account
func Register(rw http.ResponseWriter, r *http.Request) {
    // Parse request body and extract the client
    var client models.Client
    err := json.NewDecoder(r.Body).Decode(&client)
    // If there is any error in unmarshalling JSON body, return
    if err != nil {
        util.ResponseWriter(rw, http.StatusBadRequest, "Invalid JSON body", nil)
        logging.Write("[ERROR]: Error in unmarshalling JSON Client body - %v\n", err)
        return
    }

    // Validat the client fields
    errs := models.ValidateClient(&client)
    // If invalid fields, return
    if len(errs) != 0 {
        util.ResponseWriter(rw, http.StatusUnprocessableEntity, "Invalid fields", errs)
        return
    }

    // Register the client
    err = models.RegisterClient(&client)
    // If error, then determine the type and return
    if err != nil {
        fmt.Println(err)
        // If there is error in accessing database
        _, ok := err.(*util.DataBaseError)
        if ok {
            util.ResponseWriter(rw, http.StatusInternalServerError, "Can't access database", err)
            return
        }

        // If the username is not available
        _, ok = err.(*util.UsernameNotAvailableError)
        if ok {
            util.ResponseWriter(rw, http.StatusUnprocessableEntity, "Username not available", err)
            return
        }

        // If server error
        _, ok = err.(*util.ServerError)
        if ok {
            util.ResponseWriter(rw, http.StatusInternalServerError, "Internal server error", err)
            return
        }

        util.ResponseWriter(rw, http.StatusUnprocessableEntity, "Can't create new user", nil)
        return
    }

    // If the client has successfullty registered, return the response
    util.ResponseWriter(rw, http.StatusCreated, "Client registered successfully", nil)
}
