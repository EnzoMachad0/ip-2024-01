package main

import "fmt"

var n int

func main() {

	fmt.Println("digite a quantidade de numeros do array: ")
	fmt.Scan(&n)

	vetor := make([]int, n)
	for i := 0; i < len(vetor); i++ {
		fmt.Scan(&vetor[i])

	}

	fmt.Println("SAIDA")
	for i := 0; i < len(vetor); i++ {
		fmt.Printf("%d ", vetor[i])

	}
}
