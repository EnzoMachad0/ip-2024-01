package main

import "fmt"

var n, cont int

func main() {
	fmt.Println("digite o tamanho do vetor:")
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < len(vetor); i++ {

		fmt.Scan(&vetor[i])

	}

	for i := 0; i < len(vetor); i++ {

		cont = vetor[i] + cont

	}

	fmt.Print(cont)

}
