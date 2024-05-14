package main

import "fmt"

var n int

func main() {

	fmt.Println("digite a quantidade de valores que serão armazenados no vetor")
	fmt.Scan(&n)
	for n > 5000 {
		fmt.Println("digite a quantidade de valores que serão armazenados no vetor")
		fmt.Scan(&n)
	}

	arr := make([]int, n)
	fmt.Println("digite os valores do vetor: ")
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])

	}
	fmt.Println("Saída: ")
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])

	}

}
