package main

import "fmt"

func main() {
	var slice = make([]int32, 3)

	fmt.Printf("the len is %d and cap is %d \n", len(slice), cap(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Printf("address of element %d is  %p \n", i, &slice[i])
	}

	slice = append(slice, 1, 2, 3, 4)
	fmt.Printf("the len is %d and cap is %d \n", len(slice), cap(slice))
	for i := 0; i < len(slice); i++ {
		fmt.Printf("address of element %d is  %p \n", i, &slice[i])
	}
	slice = append(slice, 1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5)
	fmt.Printf("the len is %d and cap is %d \n", len(slice), cap(slice))
	for i := 0; i < len(slice); i++ {
		fmt.Printf("address of element %d is  %p \n", i, &slice[i])
	}
}
