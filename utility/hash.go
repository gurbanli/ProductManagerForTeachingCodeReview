package utility

import (
	"crypto/sha512"
	"encoding/hex"
)

//author: gurbanli

func GetHash(clearText string) string {
	h := sha512.New()
	h.Write([]byte(clearText))
	return hex.EncodeToString(h.Sum(nil))
}
