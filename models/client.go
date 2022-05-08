package models

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/killtheverse/go-chat-signal-server/db"
	"github.com/killtheverse/go-chat-signal-server/logging"
	"github.com/killtheverse/go-chat-signal-server/util"
	"go.mongodb.org/mongo-driver/bson"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
}

// Client represents a user client
type Client struct {
    // Username of the user
    Username        string      `json:"username" bson:"username" validate:"required,alphanum,min=4,max=100"`

    // Password of the user
    Password        string      `json:"password" bson:"password" validate:"required,min=8,max=100"`

    // Name of the user
    Name            string      `json:"name" bson:"name"`
    
    // Date and time the client registered for the service
    TimeCreated     time.Time   `json:"timeCreated" bson:"timeCreated"`

    // Date and Time the user last logged in
    LastLogin       time.Time   `json:"lastLogin" bson:"lastLogin"`

    // Last Refresh Token used
    RefreshToken    string      `json:"refreshToken" bson:"refreshToken"`
}

// ValidateClient validates the fields of Client
func ValidateClient(client *Client) []error {
    fieldErrors := make([]error, 0)
    err := validate.Struct(client)
    if err != nil {
        logging.Write("Errors while validating fields - %v", err)
        for _, fe := range err.(validator.ValidationErrors) {
            var errorField string
            var errorMessage string
            errorField = fe.Field()
            
            switch fe.Tag() {
                case "required":
                    errorMessage = "Field is required"
                case "min":
                    errorMessage = "Length should be atleast " + fe.Param()
                case "max":
                    errorMessage = "Length should be atmost " + fe.Param()
            }
            fieldError := util.FieldError{
                Field: errorField,
                Message: errorMessage,
            }
            fieldErrors = append(fieldErrors, &fieldError)
        }
    }
    return fieldErrors
}

// RegisterClient registers a new client
func RegisterClient(client *Client) error {
    
    // Since username should be unique, check if any user already exists with the given username
    filter := bson.M {"username": client.Username}
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    count, err := db.Database.Collection("clients").CountDocuments(ctx, filter)
    // If there is any error in processing database query, return
    if err != nil {
        logging.Write("[ERROR]: Can't process db query: %v\n", err)
        return &util.DataBaseError{Message: "Can't perform database query"}
    }
    if count > 0 {
        logging.Write("Username %v is not available\n", client.Username)
        return &util.UsernameNotAvailableError{Message: fmt.Sprintf("Username %v is not available", client.Username)}
    }

    hashedPassword, err := util.GenerateHashPassword(client.Password)
    // If theres is any error in hashing password, return it
    if err != nil {
        logging.Write("[ERROR]: Can't hash password: %v\n", err)
        return &util.ServerError{Message: "Unable to hash password"}
    }
    // Update the timeCreated and Password fields
    client.Password = hashedPassword
    client.TimeCreated = time.Now()
    
    // Insert document in database
    ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    res, err := db.Database.Collection("clients").InsertOne(ctx, client)
    if err != nil {
        logging.Write("[ERROR] Can't insert document in database: %v\n", err)         
        return &util.DataBaseError{Message: "Can't perform database operations"}
    } else {
        logging.Write("Inserted document InsertID: %v\n", res.InsertedID)
    }
    return nil
}



