package main

import "fmt"

var n int

func main() {
	fmt.Println("ENTRADA")
	fmt.Scanln(&n)

	notas := make([]int, n)

	for i := 0; i < n; i++ {

		fmt.Scanln(&notas[i])

	}
	maior, _ := theBiggest(notas)
	_, indice := theBiggest(notas)

	fmt.Println(notas)
	fmt.Println("SAÍDA")
	fmt.Printf("NOTA %d, %d VEZES \n", notas[n-1], aparicao(notas, notas[n-1]))
	fmt.Printf("NOTA %d, INDICE %d", maior, indice)

}

func aparicao(notas []int, numero int) int {

	cont := 0
	for i := 0; i < len(notas); i++ {
		if notas[i] == numero {
			cont = cont + 1

		}

	}

	return cont

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
