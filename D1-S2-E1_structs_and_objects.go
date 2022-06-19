// Including the main package
package main

// Importing fmt
import (
	"fmt"
)

type Person struct { //create a struct like Objects in OOP
	name string //define attributes of the person
}

func (p Person) greet() string { //function for Person struct
	return "Hello " + p.name + " :) "
}

func (p Person) greetSomeone(q Person) string { //function for Person struct which takes a parameter as person too
	return "Hello " + p.name + " from " + q.name + " :) "
}

func main() { //main function can only receive arguments by os.args
	var greeter Person = Person{name: "Bahadir"}             //create a new person struct with name bahadir
	var greeted Person = Person{name: "Testo"}               //create a new person to be greeted
	greeting := greeter.greet()                              //call the greet function for the struct:bahadir
	greetingToAnotherPerson := greeter.greetSomeone(greeted) //
	fmt.Println(" " + greeting)
	fmt.Println(" " + greetingToAnotherPerson)
}
