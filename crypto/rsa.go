package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
)

// 加密
func RsaEncryptS1(origData []byte, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey) ////将密钥解析成公钥实例
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	resule, err := rsa.EncryptPKCS1v15(rand.Reader, pub, origData) //RSA算法加密
	if err != nil {
		return nil, err
	}
	return resule, nil
}

// 解密
func RsaDecryptS1(cipher []byte, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey) //将密钥解析成私钥实例
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipher) //RSA算法解密
}

// 加密
func RsaEncryptS1WithString(origData string, publicKey string) (string, error) {
	data, err := RsaEncryptS1([]byte(origData), []byte(publicKey))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(data), nil
}

// 解密
func RsaDecryptS1WithString(cipher string, privateKey string) (string, error) {
	data, err := hex.DecodeString(cipher)
	if err != nil {
		return "", err
	}

	origData, err := RsaDecryptS1(data, []byte(privateKey)) //RSA解密
	if err != nil {
		return "", err
	}
	return string(origData), nil
}
