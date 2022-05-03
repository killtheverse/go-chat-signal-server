package config

import (
	"github.com/spf13/viper"
)

func init() {
    // Read the config file
    viper.SetConfigFile("DEV.env")
    viper.ReadInConfig()
}

// JWT configuration structure
type JWTConfig struct {
    
    // Lifetime of access token in minutes
    AccessTokenTL       int

    // Lifetime of refresh token in minutes
    RefreshTokenTL      int

    // Secret Key
    SecretKey           string
}

// GetJWTConfig returns the JWT configuration
func GetJWTConfig() *JWTConfig {
    config := new(JWTConfig)

    // Default value is 10 minutes
    config.AccessTokenTL = 10

    // Default value is 10 days
    config.RefreshTokenTL = 60*24*10;
    config.SecretKey = viper.GetString("SECRET_KEY")
    return config
}
