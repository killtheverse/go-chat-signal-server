package db

import (
	"context"
	"time"

	logger "github.com/killtheverse/go-chat-signal-server/logging"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    Database *mongo.Database
    DBClient *mongo.Client
)

// Connect attempts to connect to the MongoDB cluster and return the Database client
func Connect(DBURI string, DBName string) error {
    logger.Write("Connecting to Database\n")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var err error
    DBClient, err = mongo.Connect(ctx, options.Client().ApplyURI(DBURI))
    Database = DBClient.Database(DBName)
    return err
}

// Disconnect attempts to disconnect the client from the cluster
func Disconnect() error {
    logger.Write("Disconnecting from the Database\n")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err := DBClient.Disconnect(ctx)
    return err
}
