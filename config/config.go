package config

import (
	"github.com/spf13/viper"
)

func init() {
    // Read the config file
    viper.SetConfigFile("DEV.env")
    viper.ReadInConfig()
}

// Server configuration structure
type ServerConfig struct {
    
    // The address and port that the server will run on
    ServerAddress       string

    // The Database name
    DBName              string

    // The Database connection URI
    DBURI               string

}

// NewConfig creates, initializes and returns a ServerConfig instance
func NewConfig() *ServerConfig {
    serverConfig := new(ServerConfig)
    serverConfig.initialize()
    return serverConfig
}

// initialize will read the env vars and store them in config struct
func (config *ServerConfig) initialize() {
    config.ServerAddress = ":" + viper.GetString("PORT")
    config.DBName = viper.GetString("DBNAME")
    config.DBURI = viper.GetString("DBURI")
}

