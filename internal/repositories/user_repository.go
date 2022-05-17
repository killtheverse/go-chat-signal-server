package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

    return &UserRepository {
        client:         client,
        database:       client.Database(dbName),
        collection:     client.Database(dbName).Collection("clients"),
    }, nil
}

func (r* UserRepository) Login(username string, password string) error {
    // TODO: Implement login
    return nil
}

func (r* UserRepository) Register(username string, password string, name string) error {
    // TODO: Implement register
    return nil
}

