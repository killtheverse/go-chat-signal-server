package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/killtheverse/go-chat-signal-server/config"
)


var jwtConfig config.JWTConfig

func init() {
    jwtConfig = *config.GetJWTConfig()
}

// createSignedToken creates a signed token with given expiry time
func createSignedToken(username string, duration int) (string, error) {
    signingKey := []byte(jwtConfig.SecretKey)
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = username
    claims["ext"] = time.Now().Add(time.Minute*time.Duration(duration)).Unix()
    tokenString, err := token.SignedString(signingKey)
    return tokenString, err
}

// getSignedAccessToken creates and returns a Access token
func getAccessToken(username string) (string, error) {
    token, err := createSignedToken(username, jwtConfig.AccessTokenTL)
    return token, err
}

// getRefreshToken creates and returns a Refresh token
func getRefreshToken(username string) (string, error) {
    token, err := createSignedToken(username, jwtConfig.RefreshTokenTL)
    return token, err
}

