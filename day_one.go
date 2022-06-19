// Including the main package
package main

// Importing fmt
import (
	"fmt"
	"os"
)

var myInt int = 100

//yourInt := 230 //:= does declare and assign and cannot be used outside of the functions

func main() { //main function can only receive arguments by os.args
	name := os.Args[1] //here giving an argument by running go run day_one.go "testoo"

	//var name string = "Bahadir"
	var greeting = createGreet(name) //you have to have var behind any declaration
	fmt.Printf("%s", greeting)       //printf takes %s for formatting

}

func createGreet(name string) string { //a function takes a string argument "name" and returns a string
	greeting := "Hello " + name + " :) " //declaration and assignment at the same time with :=
	return greeting

}
