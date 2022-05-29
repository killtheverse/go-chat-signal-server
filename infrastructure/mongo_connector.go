package infrastructure

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/killtheverse/go-chat-signal-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
)


type MongodbConnector interface {
    DB(context.Context) *mongo.Database
    Client(context.Context) *mongo.Client
    Config() config.MongodbConfig
}

type mongodbConnector struct {
    cfg     *config.MongodbConfig
    db      *mongo.Database
    client  *mongo.Client
}

func NewMongodbConnector(cfg *config.MongodbConfig) (*mongodbConnector, error) {
    mongodbConnector := &mongodbConnector{
        cfg: cfg,
    }

    err := mongodbConnector.connect()
    if err != nil {
        return mongodbConnector, err
    }
    return mongodbConnector, nil
}

func(connector *mongodbConnector) connect() error {
    var (
        connectOnce sync.Once
        err         error
        client      *mongo.Client
    )

    connectOnce.Do(func() {
        connString := connector.cfg.URI
        client, err = mongo.NewClient(options.Client().ApplyURI(connString))
        if err != nil {
            return
        }

        ctx, cancel := context.WithTimeout(context.Background(), time.Duration(connector.cfg.DialTimeOut))
        defer cancel()
        err = client.Connect(ctx)
        if err != nil {
            return
        }
    })
    if err != nil {
        return err
    }

    connector.client = client
    connector.db = connector.client.Database(connector.cfg.DatabaseName)
    return nil
}

func (connector *mongodbConnector) DB(ctx context.Context) *mongo.Database {
    var rp readpref.ReadPref
    err := connector.client.Ping(ctx, &rp)
    if err != nil {
        fmt.Errorf("Failed to ping Database - %s\n", err)
    }
    return connector.db
}

func (connector *mongodbConnector) Client(ctx context.Context) *mongo.Client {
    return connector.client
}

func (connector *mongodbConnector) Config(ctx context.Context) *config.MongodbConfig {
    return connector.cfg
}

func (connector *mongodbConnector) EnsureIndex(collection *mongo.Collection, indexMap map[string]*options.IndexOptions) error {
    indexView := collection.Indexes()
    for k, index := range indexMap {
        if isCompositeKey(k) {
            doc := bsonx.Doc{}
            allKeys := strings.Split(k, "-")
            for _, key := range allKeys {
                elem := bsonx.Elem{key, bsonx.Int32(int32(1))}
                doc = append(doc, elem)
            }
            indexModel := mongo.IndexModel{Keys: doc, Options: index}
            _, err := indexView.CreateOne(context.Background(), indexModel)
            if err != nil {
                return err
            }
        } else {
            keys := bsonx.Doc{{Key: k, Value: bsonx.Int32(int32(1))}}
            indexModel := mongo.IndexModel{Keys: keys, Options: index}
            _, err := indexView.CreateOne(context.Background(), indexModel)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func isCompositeKey(key string) bool {
    return len(strings.Split(key, "-")) > 1
}
