package main

func reverseSlice(arr *[10]int) {

	l := len(arr)

	for i := 0; i < l/2; i++ {
		j := l - 1 - i
		arr[i], arr[j] = arr[j], arr[i]
	}

}

func reverseSlice2(arr []float64) []float64 {
	var reversedArr []float64

	for i := len(arr); i > 0; i-- {
		reversedArr = append(reversedArr, arr[i])
	}

	return reversedArr
}
