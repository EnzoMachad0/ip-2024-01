package main

import "fmt"

var n int

func main() {

	fmt.Println("ENTRADA:")
	fmt.Scanln(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])

	}

	fmt.Println("SAIDA: ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])

	}

	fmt.Println("")

	for i := (n - 1); i >= 0; i-- {
		fmt.Printf("%d ", arr[i])

	}
	fmt.Println("")

	maior, _ := theBiggest(arr)
	fmt.Println(maior)
	fmt.Println(menorElemento(arr))
}

func theBiggest(vetor []int) (maior int, indice int) {
	if len(vetor) == 0 {
		return 0, -1
	}

	maior = vetor[0]
	for i := 1; i < len(vetor); i++ {
		if vetor[i] > maior {
			maior = vetor[i]
			indice = i
		}
	}

	return maior, indice

}

func menorElemento(arr []int) int {

	menor := arr[0]
	for i := 0; i < len(arr); i++ {

		if arr[i] < menor {
			menor = arr[i]

		}
	}
	return menor

}
