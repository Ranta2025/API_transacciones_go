package utils

import "crypto/rand"

func Generate() string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijqlmnopqrstuvwxyz0123456789"

	b := make([]byte, 10)
	_, err := rand.Read(b)
	Check_err(err)

	for i := range b {
		b[i] = characters[int(b[i])% len(characters)]
	}
	return string(b)
}