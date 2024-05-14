package main

import "fmt"

var n int

func main() {

	fmt.Println("ENTRADA: ")
	fmt.Scanln(&n)

	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scanln(&arr[i])
	}

	arrSorted := bubbleSort(arr)
	fmt.Println(arr)
	fmt.Println("SAIDA")

	for i := 0; i < len(arrSorted); i++ {
		fmt.Println(arrSorted[i])

	}

}

func bubbleSort(arr []int) []int {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}
