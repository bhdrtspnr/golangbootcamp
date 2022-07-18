package main

import "fmt"

func rotateSinglePass(arr []int, rotate int) []int {
	//no idea ¯\_(ツ)_/¯
	returnArr := make([]int, len(arr))

	for i := len(arr) - 1; i > 0; i-- {
		if i < rotate { //if its less than the rotate argument it needs to be pushed to the end of the list
			returnArr[len(arr)-i] = arr[i]
		} else {
			returnArr[i-rotate] = arr[i]
		}
	}
	return returnArr

}

//test
func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	reversed := rotateSinglePass(s, 2)
	fmt.Println(reversed)
}
