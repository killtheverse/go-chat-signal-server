package config

import (
	"testing"

	"github.com/spf13/viper"
)

// TestViperReadConfig tests whether viper can read the dotenv file and checks if the environment variables are properly set
func TestViperReadConfig(t *testing.T) {
    viper.SetConfigFile("../DEV.env")
    err := viper.ReadInConfig()
    if err != nil {
        t.Fatal("Viper can't read the config file")
    }
    
    vars := []string {"PORT", "DBNAME", "DBURI"}
    for _, envVar := range vars {
        readVar := viper.GetString(envVar)
        if readVar == "" {
            t.Fatalf("The following environment variable was not read: %v", envVar)
        }
    }
}

// TestServerConfig tests whether the ServerConfig instance is properly being created or not
func TestServerConfig(t *testing.T) {
    viper.SetConfigFile("../DEV.env")
    err := viper.ReadInConfig()
    if err != nil {
        t.Fatal("Viper can't read the config file")
    }

    serverConfig := NewConfig()
    if serverConfig.ServerAddress != ":" + viper.GetString("PORT") {
        t.Fatal("Server address is not properly configured")
    }
    if serverConfig.DBName != viper.GetString("DBNAME") {
        t.Fatal("Database name is not properly configured")
    }
    if serverConfig.DBURI != viper.GetString("DBURI") {
    t.Fatal("Database URI is not properly configured")
    }
}
