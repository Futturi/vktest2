package utils

import (
	"crypto/sha1"
	"fmt"
	"os"
)

var (
	salt = os.Getenv("PASSWORD_SALT")
)

func HashPass(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
