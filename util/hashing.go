package util

import "golang.org/x/crypto/bcrypt"

// generateHashPassword generates a hash of the password string
func GenerateHashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// checkPasswordHash compares a plaintext password with it's hash value
func CheckPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
