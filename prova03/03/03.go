package main

import (
	"fmt"
)

var array []int = []int{5, 4, 3, 2, 1}

func main() {
	bubbleSort(array)
	fmt.Println(array)

}

func bubbleSort(array []int) {
	n := len(array)
	for i := 0; i < n; i++ {
		troca := false
		for j := 0; j < n-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				troca = true
			}
		}

		if !troca {
			break
		}

	}
}
