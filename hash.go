package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreateHashWithSHA256(key string) string {
	h := sha256.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}

