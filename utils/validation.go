package utils

import (
	"unicode"
)

func Validation_Upper(password string) bool {
	for _, r := range password {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func Validation_Character_special(password string) bool {
	for _, r := range password {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r){
			return true
		}
	}
	return false;
}

func Validation_Number(password string) bool {
	for _, r := range password {
		if unicode.IsNumber(r){
			return true
		}
	}
	return false;
}

func Validation_Letter(password string) bool {
	for _, r := range password {
		if unicode.IsLetter(r){
			return true
		}
	}
	return false;
}