package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/killtheverse/go-chat-signal-server/config"
)


var jwtConfig config.JWTConfig

func init() {
    jwtConfig = *config.GetJWTConfig()
}

type TokenDetails struct {
    AccessToken     string
    RefreshToken    string
    AccessUuid      string
    RefreshUuid     string
    ATExpires       int64
    RTExpires       int64
}

// CreateToken creates a TokenDetails instance
func CreateToken(username string) (*TokenDetails, error) {
    td := &TokenDetails{}
    var err error
    td.ATExpires = time.Now().Add(time.Minute*time.Duration(jwtConfig.AccessTokenTL)).Unix()
    td.AccessUuid = uuid.New().String()

    td.RTExpires = time.Now().Add(time.Minute*time.Duration(jwtConfig.RefreshTokenTL)).Unix()
    td.RefreshUuid = uuid.New().String()
    
    td.AccessToken, err = createSignedToken(username, td.ATExpires, td.AccessUuid)
    if err != nil {
        return nil, err
    }

    td.RefreshToken, err = createSignedToken(username, td.ATExpires, td.RefreshUuid)
    if err != nil {
        return nil, err
    }
    return td, nil
}

// createSignedToken creates a signed token with given expiry time
func createSignedToken(username string, expTime int64, uuid string) (string, error) {
    signingKey := []byte(jwtConfig.SecretKey)
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = username
    claims["uuid"] = uuid
    claims["exp"] = expTime
    tokenString, err := token.SignedString(signingKey)
    return tokenString, err
}

