package util

import "testing"

// TestPasswordHash tests whether passwords are hashed properly or not
func TestHashPassword(t *testing.T) {
    hash, err := generateHashPassword("password")
    if err != nil {
        t.Fatal("Can't generate password hash")
    }
    compare := checkPasswordHash("password", hash)
    if compare == false {
        t.Fatal("Password does not match with hash value")
    }
}
