package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

func Md5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

//urlencode
func Urlencode(s string) string {
	return url.QueryEscape(s)
}

//HMAC SHA256
func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	str := fmt.Sprintf("%x", h.Sum(nil))
	return strings.ToLower(str)
}

//Sha256
func Sha256(message string) string {

	h := sha256.New()
	h.Write([]byte(message))
	bs := h.Sum(nil)

	str := fmt.Sprintf("%x", bs)
	return strings.ToLower(str)
}
