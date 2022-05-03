package db

import (
	"context"
	"time"

	logger "github.com/killtheverse/go-chat-signal-server/logging"
	
        "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect attempts to connect to the MongoDB cluster and return the Database client
func Connect(DBURI string) (*mongo.Client, error) {
    logger.Write("Connecting to Database\n")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(DBURI))
    return client, err
}

// Disconnect attempts to disconnect the client from the cluster
func Disconnect(client *mongo.Client) error {
    logger.Write("Disconnecting from the Database\n")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err := client.Disconnect(ctx)
    return err
}
