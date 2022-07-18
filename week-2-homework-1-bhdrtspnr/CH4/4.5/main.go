package main

import "fmt"

func eliminate(arr []int) []int {
	returnArray := make([]int, 0) //create new slice

	for i := 0; i < len(arr); i++ { //iterate over string
		if i < len(arr)-1 { //do not enter if greater than or equal to len()-1 since it'll end up with index out of bounds error
			if arr[i] != arr[i+1] {
				returnArray = append(returnArray, arr[i]) //if n and n+1st element is not the same, add arr[i]
			}
		} else {
			returnArray = append(returnArray, arr[i]) //else to add the last element in the slice since we're not iterating over it above
		}
	}

	return returnArray
}

func main() {
	//it works
	s := []int{5, 5, 6, 6, 7, 8, 9, 9}
	c := eliminate(s)

	fmt.Println(c)
}
