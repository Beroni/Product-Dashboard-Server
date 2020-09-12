package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) string {

	password := []byte(pwd)

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func CheckPasswordHash(hashedPwd string, password string) bool {

	plainPwd := []byte(password)
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
