package utils

import "golang.org/x/crypto/bcrypt"

func Compare_password(password string, hash string) bool {
	res := bcrypt.CompareHashAndPassword([]byte (hash), []byte (password))
	return res == nil
}