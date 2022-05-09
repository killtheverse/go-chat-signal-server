package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/killtheverse/go-chat-signal-server/db"
	"github.com/killtheverse/go-chat-signal-server/logging"
	"github.com/killtheverse/go-chat-signal-server/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (c *Client) MarshalJSON() ([]byte, error) {
    var tmp struct {
        Username        string      `json:"username"`
        Name            string      `json:"name"`
        TimeCreated     time.Time   `json:"timeCreated"`
        LastLogin       time.Time   `json:"lastLogin"`

    }
    tmp.Username = c.Username
    tmp.Name = c.Name
    tmp.TimeCreated = c.TimeCreated
    tmp.LastLogin = c.LastLogin
    return json.Marshal(&tmp)
}

// ValidateClient validates the fields of Client
func ValidateClient(client *Client) []error {
    fieldErrors := make([]error, 0)
    // Validate the client structure
    err := validate.Struct(client)
    // If errors exist
    if err != nil {
        logging.Write("Errors while validating fields - %v", err)
        // Iterate through all field errors
        for _, fe := range err.(validator.ValidationErrors) {
            var errorField string
            var errorMessage string
            errorField = fe.Field()
            // Write custom messages depending on the tag
            switch fe.Tag() {
                case "required":
                    errorMessage = "Field is required"
                case "min":
                    errorMessage = "Length should be atleast " + fe.Param()
                case "max":
                    errorMessage = "Length should be atmost " + fe.Param()
            }
            // Create custom error
            fieldError := util.FieldError{
                Field: errorField,
                Message: errorMessage,
            }
            // Append custom error to fieldErrors list
            fieldErrors = append(fieldErrors, &fieldError)
        }
    }
    return fieldErrors
}

// RegisterClient registers a new client
func RegisterClient(client *Client) (*Client, error) {
    
    // Since username should be unique, check if any user already exists with the given username
    filter := bson.M {"username": client.Username}
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    err := db.Database.Collection("clients").FindOne(ctx, filter).Err()
    // If any client already exists with the same username, there will be no error
    if err == nil {
        logging.Write("Username %v is not available\n", client.Username)
        return nil, &util.UsernameNotAvailableError{Message: fmt.Sprintf("Username %v is not available", client.Username)}
    } else if err != mongo.ErrNoDocuments {
        logging.Write("[ERROR]: Can't process db query: %v\n", err)
        return nil, &util.DataBaseError{Message: "Can't perform database query"}
    }
     
    // If no existing client found with same username, continue creating client

    hashedPassword, err := util.GenerateHashPassword(client.Password)
    // If theres is any error in hashing password, return it
    if err != nil {
        logging.Write("[ERROR]: Can't hash password: %v\n", err)
        return nil, &util.ServerError{Message: "Unable to hash password"}
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
        return nil, &util.DataBaseError{Message: "Can't perform database operations"}
    } else {
        logging.Write("Inserted document InsertID: %v\n", res.InsertedID)
    }
    return client, nil
}

