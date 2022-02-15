package main

import "net/url"

func EncodeData(plainText string) string {
	return url.QueryEscape(plainText)
}
