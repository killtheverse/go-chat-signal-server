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

    // Access Token life time in minutes
    AccessTokenLT       int

    // Refresh Token life time in seconds
    RefreshTokenLT      int

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

    config.AccessTokenLT = viper.GetInt("JWT_ACCESS_LT")
    // If no value is specified in dotenv file, set the defualt value to 10 minutes
    if config.AccessTokenLT == 0 {
        config.AccessTokenLT = 10
    }
    
    config.RefreshTokenLT = viper.GetInt("JWT_REFRESH_LT")
    // If no value is specified in dotenv file, set the default value to 10 days
    if config.RefreshTokenLT == 0 {
        config.RefreshTokenLT = 10*24*60
    }
}

