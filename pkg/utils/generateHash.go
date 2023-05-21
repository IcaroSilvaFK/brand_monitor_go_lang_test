package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHash(str string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	return string(hash), err

}
