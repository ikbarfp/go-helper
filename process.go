package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func timeout(second int) {
	fmt.Print("Returning to Home\n")
	for i := second; i >= 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Printf("%v . . .\n", i)
	}
}

func subMenuInput(input int) {
	reader := bufio.NewReader(os.Stdin)
	switch input {
	case 1 :
		fmt.Println("Input your key : ")
		key, _ := reader.ReadString('\n')
		key = strings.Replace(key, "\n", "", -1)
		fmt.Println("Input your value : ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text != "" {
			encryptedValue := EncryptData(text, key)
			fmt.Printf("\nYour Encrypted Value : %+v\n\n", encryptedValue)
			//valueEncrypted := Encrypt([]byte("Hello World"), value)
			// fmt.Printf("Your Encrypted Value : %+v\n", string(valueEncrypted))
			timeout(5)
		}
	case 2 :
		fmt.Println("Input your key : ")
		key, _ := reader.ReadString('\n')
		key = strings.Replace(key, "\n", "", -1)
		fmt.Println("Input your value : ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text != "" {
			decryptedValue := DecryptData(text, key)
			fmt.Printf("\nYour Decrypted Value : %+v\n\n", decryptedValue)
			// valueDecrypted := Decrypt([]byte("Hello World"), value)
			// fmt.Printf("Your Decrypted Value : %+v\n\n", string(valueDecrypted))
			timeout(5)
		}
	case 3 :
		fmt.Println("Input your value : ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text != "" {
			encodedValue := EncodeData(text)
			fmt.Printf("\nYour Encoded Value : %+v\n\n", encodedValue)
			timeout(5)
		}
	case 4 :
		fmt.Println("Input your value : ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text != "" {
			//hashValue := CreateHashWithMD5(text)
			hashValue := CreateHashWithSHA256(text)
			fmt.Printf("\nYour Hash Value : %+v\n\n", hashValue)
			timeout(5)
		}
	case 0 :
		return
	}
}
