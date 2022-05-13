package auth

import (
	"testing"
)

// TestTokenCreate tests whether token is being created or not
func TestTokenCreate(t *testing.T) {
    _, err := CreateToken("test-user")
    if err != nil {
        t.Fatalf("Can't create token: %v", err)
    }
}
