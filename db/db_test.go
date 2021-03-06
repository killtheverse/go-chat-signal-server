package db

import (
	"testing"

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

