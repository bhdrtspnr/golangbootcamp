package main

import "fmt"

/* define global variable */
var g int

func main() {

	/* define local variables */
	var a, b int

	/* initialize variables */
	a = 10
	b = 20
	g = a + b

	fmt.Printf("Values： a = %d, b = %d and g = %d\n", a, b, g)
}
