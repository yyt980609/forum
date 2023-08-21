package md5

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "Kassadin"

// EncryptPassword 加密密码
func EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum([]byte(nil)))
}
