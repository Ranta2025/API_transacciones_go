package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte("clave-secreta")

func Generated_token(username string, rol string) (string, error){
	clans := jwt.MapClaims{
		"username" : username,
		"exp" : time.Now().Add(15 * time.Minute).Unix(),
		"rol" : rol,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clans)
	return token.SignedString(secretkey)
}