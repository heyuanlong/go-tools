package main

import (
	"encoding/base64"
	"fmt"

	"github.com/heyuanlong/go-tools/crypto"
)

func main() {

	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("sfe023f_9aaafwfl_7893345")
	result, err := crypto.AesEncrypt([]byte("1001|fdid#$%^&^*^&*fi232|1550733642"), key)
	if err != nil {
		panic(err)
	}
	lok := base64.StdEncoding.EncodeToString(result)
	fmt.Println(lok)
	nok, err := base64.StdEncoding.DecodeString(lok)
	fmt.Println(result)
	fmt.Println(nok)

	origData, err := crypto.AesDecrypt([]byte(nok), key)
	if err != nil {
		//panic(err)
	}
	fmt.Println(string(origData))
	fmt.Println(len(origData))
}
