package main

import (
	"encoding/hex"
	"fmt"

	"github.com/heyuanlong/go-tools/crypto"
)

func main() {

	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("sfe023f_9aaafwfl_7893345")
	result, err := crypto.AesEncrypt([]byte("1001|fdid#$%^&*fi232|1550733642"), key)
	if err != nil {
		panic(err)
	}
	lok := hex.EncodeToString(result)
	fmt.Println(lok)
	nok, err := hex.DecodeString(lok)
	fmt.Println(result)
	fmt.Println(nok)

	origData, err := crypto.AesDecrypt([]byte(nok), key)
	if err != nil {
		//panic(err)
	}
	fmt.Println(string(origData))
	fmt.Println()
	fmt.Println()

	result1, err1 := crypto.AesEncryptWithString("1001|fdid#$%^&*fi232|1550733642", "sfe023f_9aaafwfl_7893345")
	if err1 != nil {
		panic(err1)
	}
	origData1, err1 := crypto.AesDecryptWithString(result1, "sfe023f_9aaafwfl_7893345")
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(result1)
	fmt.Println(origData1)

}
