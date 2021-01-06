package library

import (
	"golang.org/x/crypto/bcrypt"
)

/*EncryptPassword ...
@desc Convert plainText password to ciphertext
*/
func EncryptPassword(plainText []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(plainText, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

/*CompareHashedPassword ...
@desc Compare plainText password and ciphertext
*/
func CompareHashedPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}
