package models

import (
	"testing"

	"github.com/killtheverse/go-chat-signal-server/db"
	"github.com/spf13/viper"
)

func init() {
    viper.SetConfigFile("../DEV.env")
    viper.ReadInConfig()
}

// TestClientIndex tests whether indexes are created on client collection or not
func TestClientIndex(t *testing.T) {
    _ = db.Connect(viper.GetString("DBURI"), viper.GetString("DBNAME"))
    err := CreateClientIndex()
    if err != nil {
        t.Fatalf("Can't create index on collection clients - %v", err)
    }
    _ = db.Disconnect()
}
