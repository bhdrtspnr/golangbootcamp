package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {

	var userInput string
	fmt.Println("Enter string to be hashed: ")
	fmt.Scanln(&userInput)

	fmt.Println("1 for calculating SHA256")
	fmt.Println("2 for calculating SHA384")
	fmt.Println("3 for calculating SHA512")
	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		{
			c1 := sha256.Sum256([]byte(userInput))
			fmt.Printf("hash: %x", c1)
		}
	case 2:
		{
			c1 := sha512.Sum384([]byte(userInput))
			fmt.Printf("hash: %x", c1)
		}
	case 3:
		{
			c1 := sha512.Sum512([]byte(userInput))
			fmt.Printf("hash: %x", c1)
		}
	}

}
