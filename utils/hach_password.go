package utils

import "golang.org/x/crypto/bcrypt"

func Hash_password(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte (password), bcrypt.DefaultCost)
	Check_err(err)
	string_hash := string(hash)
	return string_hash 
}