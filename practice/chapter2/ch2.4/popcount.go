package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var count int = 0
	layer := uint64(1)
	for i := 0; i < 64; i++ { //iterating over 64 times at most since the greatest available number in uint64 contains 64 bits
		if x&layer > 0 {
			count++ //if after shifting right to one bit 00010 -> 00001 -> 00000 is greater than zero, increase count by one.
		}
		x >>= 1 //shift right 1 bit
	}
	return count
}

func main() {
	fmt.Println(PopCount(0x1234567890))
}
