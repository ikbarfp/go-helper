package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// DecryptData - Using Aes Crypto Algorithm
func DecryptData(plainText, cryptorKey string) string {
	var err error
	if len(plainText) <= 0 {
		return plainText
	}
	text, err := base64.StdEncoding.DecodeString(plainText)
	block, err := aes.NewCipher([]byte(cryptorKey))
	if err != nil {
		return err.Error()
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return string(text)
}
