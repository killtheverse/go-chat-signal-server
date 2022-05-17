package domain

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
)

// User represents a user
type User struct {
    // Username of the user
    Username        string      `json:"username" bson:"username" validate:"required,alphanum,min=4,max=100"`

    // Password of the user
    Password        string      `json:"password" bson:"password" validate:"required,min=8,max=100"`

    // Name of the user
    Name            string      `json:"name" bson:"name" validate:"required,min=4,max=100"`

    // Date and time the user registered
    TimeCreated     time.Time      `json:"timeCreated" bson:"timeCreated"`

}

// NewUser creates and returns a new user
func NewUser(username string, password string, name string) *User {
    return &User{
        Username: username,                     
        Password: password,
        Name: name,
        TimeCreated: time.Now(),
    }
}

// MarshalJSON is for custom JSON Marshalling where password is not encoded
func (u *User) MarshalJSON() ([]byte, error) {
    var tmp struct {
        Username        string      `json:"username"`
        Name            string      `json:"name"`
        TimeCreated     time.Time   `json:"timeCreated"`
    }
    tmp.Username = u.Username
    tmp.Name = u.Name
    tmp.TimeCreated = u.TimeCreated
    return json.Marshal(&tmp)
}

// ValidateUser validates the fields of the user
func ValidateUser(user *User) []error {
    fieldErrors := make([]error, 0)
    
    // Validate the user structure
    validate := validator.New()
    err := validate.Struct(user)
    // If errors exist
    if err != nil {
        // Iterate through field errors
        // for _, fe := range err.(validator.ValidationErrors) {
        //     var errorField string
        //     var errorMessage string
        //     errorField = fe.Field()
        //     // Write custom messages depending on the tag
        //     switch fe.Tag() {
        //     case "required":
        //         errorMessage = "Field is required"
        //     case "min":
        //         errorMessage = "Length should be atleast " + fe.Param()
        //     case "max":
        //         errorMessage = "Length should be atmost " + fe.Param()
        //     }

        //     // Create custom error
        //     // TODO
        // }
    }
    return fieldErrors
}
