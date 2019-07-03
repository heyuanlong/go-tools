package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/heyuanlong/go-tools/crypto"
)

func main() {

	fmt.Println("aes-----------------------------------------------------------------")
	aes()

	fmt.Println("aes1-----------------------------------------------------------------")
	aes1()

	fmt.Println("rsa-----------------------------------------------------------------")
	rsa()

	fmt.Println("rsa1-----------------------------------------------------------------")
	fmt.Println()
	rsa1()
}

func aes() {
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
}

func aes1() {
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

func rsa() {

	data, err := crypto.RsaEncryptS1([]byte("polaris@studygolang.com"), publicKey) //RSA加密
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA加密:", hex.EncodeToString(data))
	fmt.Println("RSA加密:", base64.StdEncoding.EncodeToString(data))
	origData, err := crypto.RsaDecryptS1(data, privateKey) //RSA解密
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA解密:", string(origData))

}

func rsa1() {
	data, err := crypto.RsaEncryptS1WithString("polaris@studygolang.com", string(publicKey)) //RSA加密
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA加密:", data)
	origData, err := crypto.RsaDecryptS1WithString(data, string(privateKey)) //RSA解密
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA解密:", string(origData))
}

//私钥
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQClCgGZ/zRX9d108517fMh0+5xDuH3onfuTrfFKqZd5lq6Ke2jT
iwFKeGBaK9cyikRYDck2uwNcHiP4obKg6CYrzpF6y5rJouAMcZW4+HlTc9LRM/hC
TKlKu96Py+EhUe18dA9nd8wqijCbLw/wzBx6bgH+mL46gsgDdzxXaDPlgwIDAQAB
AoGAS9LQkI1Q4ZaaY5hnQmw+doyAqxZQdnZatmskX+6RorGJSCNRslr7QVkTv2nD
6TrgEmpnBuedsA1C7oBvnoB5xEhmMydzXxfIJViOQMLCPv0YgOAyzw92Z5rNHBfa
DkojPknWlPGAI5bShOBlT5k3GPaHIHT6tYlMrnYtpBITP8UCQQCvvsCs3lZ1pIC8
F64EYCVvbU/a7F7OysORKnMVPfr3I4ysGuxXU8du6rjvzx15Eh/3RVImY6QJbm/q
X0R9aVcnAkEA8GetqbS86nQZAt72Sd+uTAYx2YB6fX7Vg4VhUrRKDMaTGd8bqHa+
1GWaqrzaU2eFIsNmKVoI9i/0lGZljJZYRQJAMDJa+s2a3nZ3y52e3ppTieRrkvlh
4speqc//caLm0aIRMR3NFQHn3rZGc5XUWmCrHZAIQHjxApkj3h20VcRu3wJASdlP
e6ZNsiff1wXu2lqgDDKK9amF9y8TH8fFUcaYSLxnS7dBo8p2bICZtoE1ABH4z+j+
ZQ2HWzj4BO4/m6RDkQJAdWnyNv7rCObUr71UDzpcw8by8bhl3kVZ0loU2TYsZAFT
XtLuaVH3vSSMwggt7Y2JsFIpm+lnuAW/TmO6m3p3/g==
-----END RSA PRIVATE KEY-----
`)

//公钥
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQClCgGZ/zRX9d108517fMh0+5xD
uH3onfuTrfFKqZd5lq6Ke2jTiwFKeGBaK9cyikRYDck2uwNcHiP4obKg6CYrzpF6
y5rJouAMcZW4+HlTc9LRM/hCTKlKu96Py+EhUe18dA9nd8wqijCbLw/wzBx6bgH+
mL46gsgDdzxXaDPlgwIDAQAB
-----END PUBLIC KEY-----
`)
