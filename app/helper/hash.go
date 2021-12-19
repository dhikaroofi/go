package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ValidateMd5(hash, text string) bool {
	if Md5Hash(text) == hash {
		return true
	}
	return false
}
