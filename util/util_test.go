package util

import "testing"

// TestPasswordHash tests whether passwords are hashed properly or not
func TestHashPassword(t *testing.T) {
    hash, err := GenerateHashPassword("password")
    if err != nil {
        t.Fatal("Can't generate password hash")
    }
    compare := CheckPasswordHash("password", hash)
    if compare == false {
        t.Fatal("Password does not match with hash value")
    }
}
