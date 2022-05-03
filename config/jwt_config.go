package config

// JWT configuration structure
type JWTConfig struct {
    
    // Lifetime of access token in minutes
    AccessTokenTL       int

    // Lifetime of refresh token in minutes
    RefreshTokenTL      int
}

// GetJWTConfig returns the JWT configuration
func GetJWTConfig() *JWTConfig {
    config := new(JWTConfig)

    // Default value is 10 minutes
    config.AccessTokenTL = 10

    // Default value is 10 days
    config.RefreshTokenTL = 60*24*10;
    return config
}
