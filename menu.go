package main

import "fmt"

func menuInput() int {
	var input int
	fmt.Println("\033[H\033[2J")
	fmt.Print("1. Encrypt\n" +
		"2. Decrypt\n" +
		"3. Encode\n" +
		"4. Hashing\n")
	fmt.Println("Choose your number : ")
	fmt.Scanf("%d", &input)
	return input
}
