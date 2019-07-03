package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
)

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) (b []byte, rerr error) {
	defer func() {
		if err := recover(); err != nil {
			rerr = errors.New(fmt.Sprintf("AesEncrypt fail:%v", err))
		}
	}()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pKCS5Padding(origData, blockSize)
	// origData = zeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	return crypted, nil
}

func AesDecrypt(crypted, key []byte) (b []byte, rerr error) {
	defer func() {
		if err := recover(); err != nil {
			rerr = errors.New(fmt.Sprintf("AesDecrypt fail:%v", err))
		}
	}()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pKCS5UnPadding(origData)
	// origData = zeroUnPadding(origData)

	return origData, nil
}

func AesEncryptWithString(origData string, key string) (b string, rerr error) {
	result, err := AesEncrypt([]byte(origData), []byte(key))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(result), nil
}

func AesDecryptWithString(crypted string, key string) (b string, rerr error) {
	hexCrypted, err := hex.DecodeString(crypted)
	if err != nil {
		return "", err
	}

	origData, err := AesDecrypt(hexCrypted, []byte(key))
	if err != nil {
		return "", err
	}

	return string(origData), nil
}
