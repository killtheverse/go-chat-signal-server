package config

import (
    "os"
)

// Server configuration structure
type ServerConfig struct {
    
    // The address and port that the server will run on
    ServerAddress       string

    // The Database name
    DBName              string

    // The Database connection URI
    DBURI  string

}

// NewConfig creates, initializes and returns a ServerConfig instance
func NewConfig() *ServerConfig {
    serverConfig = new(ServerConfig)
    ServerConfig.initialize()
    return serverConfig
}

// initialize will read the env vars and store them in config struct
func (config *ServerConfig) initialize() {
    config.ServerAddress = ":" + os.Getenv("PORT")
    config.DBName = os.Getenv("DBNAME")
    config.DBURI = os.Getenv("DBURI")
}

