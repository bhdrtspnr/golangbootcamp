package main

import "fmt"

func main() {
	const (
		january = iota
		february
		march
		april
		may
		june
		july
	)

	fmt.Println(january)
	fmt.Println(june)
}
