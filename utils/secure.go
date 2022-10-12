package utils

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(pwd string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

func ComparePassword(pwd string, input string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(input))
	if err != nil {
		return false
	}

	return true
}
