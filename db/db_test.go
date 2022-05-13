package db

import (
	"context"
	"testing"
	"time"

	"github.com/spf13/viper"
)

// TestDBConnection attempts to connect and then disconnect from the remote database cluster
func TestDBConnection(t *testing.T) {
    viper.SetConfigFile("../DEV.env")
    viper.ReadInConfig()
    DBURI := viper.GetString("DBURI")
    DBName := viper.GetString("DBNAME")
    err := Connect(DBURI, DBName)
    if err != nil {
        t.Fatalf("Can't connect to database: %v", err)
    }

    err = Disconnect()
    if err != nil {
        t.Fatalf("Can't disconnect from the database: %v", err)
    }
}

//TestRedisConnection attempts to connect to redis 
func TestRedisConnection(t *testing.T) {
    viper.SetConfigFile("../DEV.env")
    viper.ReadInConfig()
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
    defer cancel()
    _, err := RedisClient.Ping(ctx).Result()
    if err != nil {
        t.Fatalf("Can't connect to redis: %v", err)
    }
}
