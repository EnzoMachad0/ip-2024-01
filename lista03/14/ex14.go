package main

import (
	"fmt"
	"math"
)

var n int

func main() {
	fmt.Println("ENTRADA: ")
	fmt.Scanln(&n)

	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scanln(&arr[i])
	}

	tamanhoDoArr := len(arr)
	arrSorted := bubbleSort(arr)
	meioImpar := int(math.Round(float64(tamanhoDoArr) / 2))
	meioPar1 := tamanhoDoArr/2 - 1
	meioPar2 := tamanhoDoArr / 2

	var mediana float64
	if tamanhoDoArr%2 != 0 {
		mediana = float64(arrSorted[meioImpar])
	} else {
		mediana = (float64(arrSorted[meioPar1]) + float64(arrSorted[meioPar2])) / 2
	}

	fmt.Println("SAIDA: ")
	fmt.Printf("%.2f \n", mediana)

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
