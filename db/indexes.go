package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// SetIndexes accepts a collection and keys as parameters and creates a unique index in the collection on the provided keys
func SetIndexes(collection *mongo.Collection, keys bsonx.Doc) error {
    index := mongo.IndexModel{}
    index.Keys = keys
    unique := true
    index.Options = &options.IndexOptions{
        Unique: &unique,
    }
    opts := options.CreateIndexes().SetMaxTime(10*time.Second)
    _, err := collection.Indexes().CreateOne(context.Background(), index, opts)
    return err
}
