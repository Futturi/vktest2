package utils

import (
	"crypto/sha1"
	"fmt"
)

const (
	salt = "woregjflqwjeiwqjmfljgvnkrjweoiqopoqghewkasdcknvbjnrwekuwiejlqwnb"
)

func HashPass(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
