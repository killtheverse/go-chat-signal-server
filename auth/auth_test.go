package auth

import (
	"testing"
)

// TestTokenCreate checks whether tokens are being created or not through getAccessToken and getRefreshToken
func TestTokenCreate(t *testing.T) {
    _, err := getAccessToken("test-user")
    if err != nil {
        t.Fatalf("Can't create access token: %v", err)
    }
    _, err = getRefreshToken("test-user")
    if err != nil {
        t.Fatalf("Can't create refresh token: %v", err)
    }
}
