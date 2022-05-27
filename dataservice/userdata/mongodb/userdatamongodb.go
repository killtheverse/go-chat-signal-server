package mongodb

import (
	"context"
	"time"

	"github.com/killtheverse/go-chat-signal-server/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
    MongoDBTimeOut = 5
)

type UserDataMongoDB struct {
    client *mongo.Client
    database *mongo.Database
    collection *mongo.Collection
}

func (udm *UserDataMongoDB) Find(username string) (*model.User, error) {
    return nil, nil
}

// Insert inserts the user in the collection
func (udm *UserDataMongoDB) Insert(user *model.User) (*model.User, error) {
    // Hash the password
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    if err != nil {
        return nil, err
    }

    // Update the ID, time created and password
    user.ID = primitive.NewObjectID()
    user.Password = string(hashedBytes)
    user.TimeCreated = time.Now()

    // Insert in the database
    ctx, cancel := context.WithTimeout(context.Background(), MongoDBTimeOut*time.Second)
    defer cancel()
    _, err = udm.collection.InsertOne(ctx, user)
    if err != nil {
        return nil, err
    }
    return user, nil
}
