package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/killtheverse/go-chat-signal-server/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
)

const (
    MongoClientTimeout = 5
)

type UserRepository struct {
    client      *mongo.Client
    database    *mongo.Database
    collection  *mongo.Collection
}

func NewUserRepository(conn string, dbName string) (*UserRepository, error) {
    ctx, cancel := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)                
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(
        conn,
    ))
    if err != nil {
        return nil, err
    }
    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        return nil, err
    }

    database := client.Database(dbName)
    collection := database.Collection("users")
    keys := bson.M {
        "username": 1,
    }
    unique := true
    indexOptions := options.IndexOptions {
        Unique: &unique,
    }
    indexModel := mongo.IndexModel {
        Keys: keys,
        Options: &indexOptions,
    }
    opts := options.CreateIndexes().SetMaxTime(MongoClientTimeout*time.Second)
    _, err = collection.Indexes().CreateOne(context.Background(), indexModel, opts)
    if err != nil {
        return nil, err
    }

    return &UserRepository {
        client:         client,
        database:       database,
        collection:     collection,
    }, nil
}

func (r* UserRepository) Login(username string, password string) error {
    // TODO: Implement login
    return nil
}

func (r* UserRepository) Register(username string, password string, name string) error {
    // Since the username should be unique, check if any user already exists with the given username
    filter := bson.M {"username": username}
    ctx, cancel := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
    defer cancel()

    err := r.collection.FindOne(ctx, filter).Err()
    // If any user already exists with the same username, there will be no error
    if err == nil {
        return errors.New("Username not available")
    } else if err != mongo.ErrNoDocuments {
        return errors.New("Can't access database")
    } 

    // If no existing user found with the same username, continue creating the user
    
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    // If there is any error in hashing password, return it 
    if err != nil {
        return err
    }
    hashedPassword := string(hashedBytes)

    user := domain.NewUser(username, hashedPassword, name)

    // Insert document in database
    ctx, cancel = context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
    defer cancel()
    _, err = r.collection.InsertOne(ctx, user)
    if err != nil {
        return err
    }
    return nil
}

