package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"
)

// HashPassword HashPassword
func HashPassword(pass, salt string) string {
	h := sha1.New()
	h.Write([]byte(pass + salt))
	checksum := h.Sum(nil)
	passHashed := base64.StdEncoding.EncodeToString([]byte(string(checksum) + salt))
	return passHashed
}

// Md5 Md5Sign
func Md5(src string) string {
	h := md5.New()
	io.WriteString(h, src)
	return hex.EncodeToString(h.Sum(nil))
}
