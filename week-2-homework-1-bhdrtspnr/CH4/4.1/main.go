package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(compareHash(c1, c2))

}

func compareHash(hash1 [32]byte, hash2 [32]byte) int {

	count := 0
	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			count++
		}
	}
	return count

}
