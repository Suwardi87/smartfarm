package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword untuk mengenkripsi password sebelum disimpan
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// CheckPassword untuk membandingkan password input dengan hash di DB
func CheckPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
