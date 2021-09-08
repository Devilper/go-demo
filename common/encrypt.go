package common

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(passwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	passwd = string(hash)
	return passwd
}

func DecryptPassword(passwd string, loginPasswd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(passwd), []byte(loginPasswd)); err != nil {
		return false
	} else {
		return true
	}
}
