package auth

import (
	"github.com/killtheverse/go-chat-signal-server/models"
	"github.com/killtheverse/go-chat-signal-server/util"
	"go.mongodb.org/mongo-driver/mongo"
)

// Credentials represents the form of credentials that a user enters to authenicate themselves
type Credentials struct {
    Username        string      `json:"username"`
    Password        string      `json:"password"`
}

func Authenticate(creds *Credentials) (string, string, error) {
    // Get the corresponding client for the username
    client, err := models.GetClient(creds.Username)
    // If any error in retreiving matching client
    if err != nil {
        // If no client found with the given username
        if err == mongo.ErrNoDocuments {
            return "", "", &util.InvalidCredentialsError{}
        } else {
            return "", "", &util.ServerError{}
        }
    }

    // Check the password with client's hashed password
    check := util.CheckPasswordHash(creds.Password, client.Password)
    // If the password does not match with hash, report error
    if !check {
        return "", "", &util.InvalidCredentialsError{}
    }

    return "", "", nil 
}

