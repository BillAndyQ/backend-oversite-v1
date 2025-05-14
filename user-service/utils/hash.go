package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    return string(bytes), err
}

func CheckPassword(rawPassword, hashed string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(rawPassword))
    return err == nil
}
