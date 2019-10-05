package gu

import (
	"crypto/md5"
	"encoding/hex"
)

// return the str's md5
func ToMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
