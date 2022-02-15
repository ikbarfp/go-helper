package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// EncryptData - Using Aes Crypto Algorithm
func EncryptData(plainText, cryptorKey string) string {
	var err error
	if len(plainText) <= 0 {
		return plainText
	}
	text := []byte(plainText)
	block, err := aes.NewCipher([]byte(cryptorKey))
	if err != nil {
		return err.Error()
	}
	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return err.Error()
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], text)
	return base64.StdEncoding.EncodeToString(cipherText)
}
