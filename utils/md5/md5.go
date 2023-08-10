package md5

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "Kassadin"

func EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
