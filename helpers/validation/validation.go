package validation

import (
	"golang.org/x/crypto/bcrypt"
)

//CheckAuthHash - validate encrypt Authorization
func CheckAuthHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
