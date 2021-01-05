package library

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(plainText []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(plainText, bcrypt.MinCost)
	
	if err != nil {
        return "", err
	}
	
	return string(hash), nil
}